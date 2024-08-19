package admin

import (
	"delivery/configs"
	"delivery/logger"
	"delivery/storage"
)

type AdminController interface {

}

type adminController struct {
	log logger.LoggerI
	storage storage.Storage
	cfg *configs.Configuration
}

func NewAdminController(log logger.LoggerI, storage storage.Storage) AdminController {
	return adminController{
		log: log,
		storage: storage,
		cfg: configs.Config(),
	}
}

