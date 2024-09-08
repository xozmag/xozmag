package entities

import (
	"delivery/pkg/utils"
	"errors"
	"time"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SendCodeReq struct {
	PhoneNumber string `json:"phone" validate:"required,phone"`
}

func (req *SendCodeReq) Validate() error {
	if !utils.IsPhoneValid(req.PhoneNumber) {
		return errors.New("invalid phone number: must be in format +99XXXXXXXXXX")
	}
	return nil
}

type RegistrReq struct {
	ID          string
	PhoneNumber string `json:"phone" validate:"required,phone" gorm:"type:varchar(13);not null;unique;"`
	Code        string `json:"code" validate:"required,len=6" gorm:"-"`
	FcmToken    string `json:"fcm_token" gorm:"fcm_token"`
}

func (req *RegistrReq) Validate() error {
	if !utils.IsPhoneValid(req.PhoneNumber) {
		return errors.New("invalid phone number: must be in format +99XXXXXXXXXX")
	}
	if len(req.Code) != 6 {
		return errors.New("invalid code: must be 6 characters long")
	}
	return nil
}

type RegistrRes struct {
	ID     string `json:"id"`
	Tokens Tokens `json:"tokens"`
}

type UserProfile struct {
	ID         string    `json:"id" gorm:"type:uuid;primaryKey"`
	Firstname  string    `json:"firstname" gorm:"column:firstname"`
	Surname    string    `json:"surname" gorm:"column:surname"`
	PhoneNumber string   `json:"phone" gorm:"column:phone_number"`
	Middlename string    `json:"middlename" gorm:"column:fathersname"`
	Birthdate  time.Time `json:"birthdate" gorm:"column:birthdate"`
	Gender     string    `json:"gender" gorm:"column:gender"`
	CreatedBy  string    `json:"created_by" gorm:"column:created_by"`
	UpdatedBy  string    `json:"updated_by" gorm:"column:updated_by"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
	//DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type UserLocation struct {
	ID string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID    string  `gorm:"user_id" json:"user_id"`
	Name      string  `gorm:"name" json:"name"`
	Latitude  float64 `gorm:"latitude" json:"latitude"`
	Longitude float64 `gorm:"longitude" json:"longitude"`
}
