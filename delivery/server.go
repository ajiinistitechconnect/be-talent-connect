package delivery

import (
	"fmt"

	"github.com/alwinihza/talent-connect-be/config"
	"github.com/alwinihza/talent-connect-be/delivery/controller"
	"github.com/alwinihza/talent-connect-be/manager"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/utils/authenticator"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Server struct {
	ucManager    manager.UsecaseManager
	engine       *gin.Engine
	host         string
	tokenService authenticator.AccessToken
}

func (s *Server) initController() {
	controller.NewRoleController(s.engine, s.ucManager.RoleUc())
	controller.NewUserController(s.engine, s.ucManager.UserUc())
	controller.NewAuthController(s.engine, s.ucManager.AuthUc(), s.tokenService)
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

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.RedisConfig.Password,
		DB:       cfg.Db,
		Username: "username",
	})

	tokenService := authenticator.NewTokenService(*cfg, client)

	repo := manager.NewRepoManager(infra)
	uc := manager.NewUsecaseManager(repo, cfg)

	r := gin.Default()
	r.GET("/migration", func(ctx *gin.Context) {
		infra.Migrate(
			&model.User{},
			&model.Role{},
		)
	})

	return &Server{
		ucManager:    uc,
		engine:       r,
		host:         fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort),
		tokenService: tokenService,
	}
}
