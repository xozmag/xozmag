package admin

import (
	"context"
	"delivery/configs"
	"delivery/entities"
	"delivery/logger"
	"delivery/storage"
	"fmt"

	"go.uber.org/zap"
)

type AdminController interface {
	CreateXozmak(ctx context.Context, req entities.Xozmak) error
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

func (a adminController) CreateXozmak(ctx context.Context, req entities.Xozmak) error{
	a.log.Info("CreateXozmak started: ", 
		zap.String("Request: ", fmt.Sprintf("XozmakID: %s, XozmakName: %s, CreatedBy: %s", req.ID, req.Name, req.CreatedBy)))

	
	err := a.storage.Admin().CreateXozmak(ctx, req)
	if err != nil {
		return err
	}

	a.log.Info("CreateXozmak finished")

	return nil
}