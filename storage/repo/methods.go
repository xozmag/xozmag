package repo

import (
	"context"
	"delivery/entities"
)

// IAdminStorage account storage interface
type IAdminStorage interface {
	CreateXozmak(ctx context.Context, req entities.Xozmak) (error)
}