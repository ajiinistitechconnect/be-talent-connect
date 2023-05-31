package delivery

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/config"
	"github.com/alwinihza/talent-connect-be/delivery/controller"
	"github.com/alwinihza/talent-connect-be/manager"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ucManager manager.UsecaseManager
	engine    *gin.Engine
	host      string
}

func (s *Server) initController() {
	controller.NewRoleController(s.engine, s.ucManager.RoleUc())
	controller.NewUserController(s.engine, s.ucManager.UserUc())
	controller.NewMentoringScheduleController(s.engine, s.ucManager.MentoringScheduleUsecase())
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	infra, err := manager.NewInfraManager(cfg)
	if err != nil {
		panic(err)
	}

	repo := manager.NewRepoManager(infra)
	uc := manager.NewUsecaseManager(repo)

	r := gin.Default()
	r.GET("/migration", func(ctx *gin.Context) {
		infra.Migrate(
			&model.User{},
			&model.Role{},
			&model.MentorMentee{},
			&model.MentoringSchedule{},
		)
	})

	return &Server{
		ucManager: uc,
		engine:    r,
		host:      fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort),
	}
}
