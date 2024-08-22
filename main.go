package main

import (
	"delivery/configs"
	"delivery/constants"
	admincontroller "delivery/controllers/admin"
	"delivery/handlers"
	"delivery/logger"
	"delivery/routers"
	"delivery/storage"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

	redisClient := redis.NewClient(&redis.Options{
        Addr: configs.Config().RedisAddr,
        Password: configs.Config().RedisPassword, // parol kerak bo'lsa
        DB: 0, // Redis DB nomeri, default 0
    })

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
	router := routers.New(h, cfg, log)

	router.Start()

}
