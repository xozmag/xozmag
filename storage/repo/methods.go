package repo

import (
	"context"
	"delivery/entities"
)

// IAdminStorage account storage interface
type IAdminStorage interface {
	CreateXozmak(ctx context.Context, req entities.Xozmak) error
	Signup(ctx context.Context, req entities.VerifyCodeReq) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (entities.LoginPostgres, error)
	UpdateUser(ctx context.Context, userID string, updateData map[string]interface{}) error
}
