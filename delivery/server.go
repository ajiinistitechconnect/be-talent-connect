package delivery

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/config"
	"github.com/alwinihza/talent-connect-be/delivery/controller"
	"github.com/alwinihza/talent-connect-be/delivery/middleware"
	"github.com/alwinihza/talent-connect-be/manager"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/utils/authenticator"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Server struct {
	ucManager    manager.UsecaseManager
	engine       *gin.Engine
	authRoute    gin.IRoutes
	host         string
	tokenService authenticator.AccessToken
}

func (s *Server) initController() {
	controller.NewRoleController(s.engine, s.ucManager.RoleUc())
	controller.NewUserController(s.engine, s.ucManager.UserUc())
	controller.NewMentoringScheduleController(s.engine, s.ucManager.MentoringScheduleUc())
	controller.NewMentorMenteeController(s.engine, s.ucManager.MentorMenteeUc())
	controller.NewProgramController(s.engine, s.authRoute, s.ucManager.ProgramUc(), s.ucManager.UserUc())
	controller.NewActivityController(s.engine, s.ucManager.ActivityUc())
	controller.NewParticipantController(s.engine, s.ucManager.ParticipantUc())
	controller.NewQuestionController(s.engine, s.ucManager.QuestionUc())
	controller.NewQuestionCategoryController(s.engine, s.ucManager.QuestionCategoryUc())
	controller.NewEvaluationCategoryController(s.engine, s.ucManager.EvaluationCategoryUc())
	controller.NewAuthController(s.engine, s.ucManager.AuthUc(), s.tokenService)
	controller.NewEvaluationController(s.engine, s.authRoute, s.ucManager.EvaluationUc(), s.ucManager.UserUc())
	controller.NewQuestionAnswerController(s.engine, s.ucManager.QuestionAnswerUc())
}

func (s *Server) Run() {
	s.initController()

	s.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://talent-connect-dev.netlify.app"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
	}))
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

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.RedisConfig.Password,
		DB:       cfg.Db,
		// Username: "username",
	})

	tokenService := authenticator.NewTokenService(*cfg, client)

	repo := manager.NewRepoManager(infra)
	uc := manager.NewUsecaseManager(repo, cfg)

	r := gin.Default()
	r.GET("/migration", func(ctx *gin.Context) {
		infra.Migrate(
			&model.User{},
			&model.Role{},
			&model.Answer{},
			&model.Question{},
			&model.Option{},
			&model.QuestionAnswer{},
			&model.QuestionCategory{},
			&model.EvaluationCategoryQuestion{},
			&model.Program{},
			&model.Evaluation{},
			&model.MentorMentee{},
			&model.MentoringSchedule{},
			&model.Activity{},
			&model.Participant{},
		)
	})

	auth := r.Group("/auth").Use(middleware.NewTokenValidator(tokenService).RequireToken())

	return &Server{
		ucManager:    uc,
		engine:       r,
		authRoute:    auth,
		host:         fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort),
		tokenService: tokenService,
	}
}
