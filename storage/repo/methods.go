package repo

import (
	"context"
	"delivery/entities"
)

// IAdminStorage account storage interface
type IAdminStorage interface {
	CreateXozmak(ctx context.Context, req entities.Xozmak) error
	Signup(ctx context.Context, req entities.SignupReq) error
	GetUserByPhone(ctx context.Context, phoneNumber string) (entities.LoginPostgres, error)
}
