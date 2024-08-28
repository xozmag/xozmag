package main

import (
	"delivery/configs"
	"delivery/constants"
	admincontroller "delivery/controllers/admin"
	"delivery/handlers"
	"delivery/logger"
	"delivery/middlewares"
	pkgutil "delivery/pkg/utils"
	"delivery/routers"
	"delivery/storage"

	"github.com/gin-gonic/gin"
)

func main() {

	//configuration settings
	cfg := configs.Config()

	// take environment from config then set gin mode according to it
	switch cfg.Environment {
	case constants.DebugMode:
		gin.SetMode(gin.DebugMode)
	case constants.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	//logger
	log := logger.NewLogger(cfg.AppName, cfg.LogLevel)
	defer logger.Cleanup(log)

	//storage init
	strg := storage.New(cfg)

	redisClient := pkgutil.NewRedisClient(*cfg)

	//controllers init
	admincontroller := admincontroller.NewAdminController(log, strg, redisClient)

	//handlers init
	h := handlers.New(
		cfg,
		log,
		admincontroller,
		redisClient,
	)

	//routers
	router := routers.New(h, cfg, log, &middlewares.JWTRoleAuthorizer{})

	router.Start()

}
