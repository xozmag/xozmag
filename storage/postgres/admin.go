package postgres

import (
	"context"
	"delivery/constants"
	"delivery/entities"
	e "delivery/errors"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type adminRepo struct {
	db *gorm.DB
}

func NewAdmin(db *gorm.DB) *adminRepo {
	return &adminRepo{db: db}
}

func (a adminRepo) CreateXozmak(ctx context.Context, req entities.Xozmak) error {
	res := a.db.WithContext(ctx).Table("xozmaks").Create(&req)
	if res.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(res.Error, &pgErr) && pgErr.Code == constants.PGUniqueKeyViolationCode {
			return fmt.Errorf("error in CreateXozmak: %w", constants.ErrXozmakAlreadyExists)
		}
		return fmt.Errorf("error in CreateXozmak: %w", res.Error)
	}
	if res.RowsAffected == 0 {
		return fmt.Errorf("error in CreateXozmak: %w", constants.ErrRowsAffectedIsZero)
	}
	return nil
}

func (a adminRepo) Signup(ctx context.Context, req entities.VerifyCodeReq) error {
	res := a.db.WithContext(ctx).Table("users").Create(req)
	if res.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(res.Error, &pgErr) && pgErr.Code == constants.PGUniqueKeyViolationCode {
			return e.ErrAccountAlreadyExists
		}
		return fmt.Errorf("error in Signup: %w", res.Error)
	}

	return nil
}

func (a adminRepo) GetUserByPhone(ctx context.Context, phoneNumber string) (entities.LoginPostgres, error) {
	user := entities.LoginPostgres{}

	res := a.db.WithContext(ctx).Table("users").Where("phone_number = ?", phoneNumber).First(&user)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return entities.LoginPostgres{}, fmt.Errorf("no any user found with phone number %s", phoneNumber)
		}
		return entities.LoginPostgres{}, fmt.Errorf("error in GetUserByPhone: %w", res.Error)
	}

	if res.RowsAffected == 0 {
		return entities.LoginPostgres{}, fmt.Errorf("no user found with phone number %s", phoneNumber)
	}

	return user, nil
}

func (a *adminRepo) UpdateUser(ctx context.Context, userID string, updateData map[string]interface{}) error {
    // Update profile data for the user with the given userID
    if err := a.db.WithContext(ctx).Table("users").Where("id = ?", userID).Updates(updateData).Error; err != nil {
        return fmt.Errorf("failed to update profile data: %w", err)
    }

    return nil
}
