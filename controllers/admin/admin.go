package admin

import (
	"context"
	"delivery/configs"
	"delivery/constants"
	"delivery/entities"
	"delivery/logger"
	pkgerrors "delivery/pkg/errors"
	"delivery/pkg/jwt"
	"delivery/storage"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AdminController interface {
	Login(ctx context.Context, req entities.LoginReq) (entities.LoginRes, error)
	VerifyCode(ctx context.Context, req entities.VerifyCodeReq) (bool, error)
	CreateXozmak(ctx context.Context, req entities.Xozmak) error
}

type adminController struct {
	log     logger.LoggerI
	storage storage.Storage
	cfg     *configs.Configuration
	redis   *redis.Client
}

func NewAdminController(log logger.LoggerI, storage storage.Storage, redis *redis.Client) AdminController {
	return adminController{
		log:     log,
		storage: storage,
		cfg:     configs.Config(),
		redis:   redis,
	}
}

func (a adminController) Login(ctx context.Context, req entities.LoginReq) (entities.LoginRes, error) {

	admin, err := a.storage.Admin().GetUserByPhone(ctx, req.PhoneNumber)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.LoginRes{}, fmt.Errorf("user with phone number %s not found", req.PhoneNumber)
		}
		a.log.Error("error fetching admin", logger.Error(err))
		return entities.LoginRes{}, err
	}

	tokenMetadata := map[string]string{
		"id":          admin.Id,
		"phoneNumber": admin.PhoneNumber,
		"role":        constants.UserRole,
	}

	tokens := entities.Tokens{}
	tokens.AccessToken, err = jwt.GenerateNewJWTToken(tokenMetadata, constants.JWTAccessTokenExpireDuration, a.cfg.JWTSecretKey)
	if err != nil {
		a.log.Error("calling GenerateNewTokens failed", logger.Error(err))
		return entities.LoginRes{}, err
	}

	tokens.RefreshToken, err = jwt.GenerateNewJWTToken(tokenMetadata, constants.JWTRefreshTokenExpireDuration, a.cfg.JWTSecretKey)
	if err != nil {
		a.log.Error("calling GenerateNewTokens failed", logger.Error(err))
		return entities.LoginRes{}, err
	}

	return entities.LoginRes{
		ID:     admin.Id,
		Tokens: tokens,
	}, nil
}

func (a adminController) VerifyCode(ctx context.Context, req entities.VerifyCodeReq) (bool, error) {
	// Redisdan saqlangan kodni olish
	redisKey := fmt.Sprintf("verification_code:%s", req.PhoneNumber)
	fmt.Println(redisKey)
	storedCode, err := a.redis.Get(ctx, redisKey).Result()
	if err == redis.Nil {
		return false, pkgerrors.NewError(http.StatusBadRequest, "Code is invalid or expired")
	} else if err != nil {
		a.log.Error("Redisdan kodni olishda xatolik", logger.Error(err))
		return false, pkgerrors.NewError(http.StatusInternalServerError, "Kod tekshirishda xatolik")
	}

	// Kiritilgan kodni tekshirish
	if storedCode != req.Code {
		return false, pkgerrors.NewError(http.StatusBadRequest, "Kod noto'g'ri")
	}

	err = a.storage.Admin().Signup(ctx, entities.SignupReq{
		ID:          uuid.NewString(),
		PhoneNumber: req.PhoneNumber,
	})
	if err != nil {
		a.log.Error("Telefon raqamini saqlashda xatolik", logger.Error(err))
		return false, pkgerrors.NewError(http.StatusInternalServerError, "Telefon raqamini saqlashda xatolik")
	}

	return true, nil
}

func (a adminController) CreateXozmak(ctx context.Context, req entities.Xozmak) error {
	a.log.Info("CreateXozmak started: ",
		zap.String("Request: ", fmt.Sprintf("XozmakID: %s, XozmakName: %s, CreatedBy: %s", req.ID, req.Name, req.CreatedBy)))

	err := a.storage.Admin().CreateXozmak(ctx, req)
	if err != nil {
		return err
	}

	a.log.Info("CreateXozmak finished")

	return nil
}
