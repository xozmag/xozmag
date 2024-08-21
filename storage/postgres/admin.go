package postgres

import (
	"context"
	"delivery/entities"
	"delivery/constants"

	"errors"
	"fmt"

	"github.com/jackc/pgconn"
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
