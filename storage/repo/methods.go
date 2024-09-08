package repo

import (
	"context"
	"delivery/entities"
)

// IAdminStorage account storage interface
type IAdminStorage interface {
	CreateXozmak(ctx context.Context, req entities.Xozmak) error
	Registration(ctx context.Context, req entities.RegistrReq) error
	UpdateUser(ctx context.Context, userID string, updateData entities.UserProfile) error
	InsertUserLocation(ctx context.Context, req entities.UserLocation) error
	GetUserProfile(ctx context.Context, id string)(entities.UserProfile, error)
	GetUserLocation(ctx context.Context, userId string)([]entities.UserLocation, error)
	GetXozmak(ctx context.Context)([]entities.Xozmak, error)
	UpdateXozmak(ctx context.Context, req entities.Xozmak) error
	DeleteXozmak(ctx context.Context, id string) error
}
