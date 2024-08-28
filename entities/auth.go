package entities

import (
	"delivery/pkg/utils"
	"errors"
	"time"
)

type LoginReq struct {
	PhoneNumber string `json:"phone" validate:"required,phone"`
}

func (req *LoginReq) Validate() error {
	if !utils.IsPhoneValid(req.PhoneNumber) {
		return errors.New("invalid phone number")
	}
	return nil
}

type LoginRes struct {
	ID     string `json:"id"`
	Tokens Tokens `json:"tokens"`
}

type SignupRes struct {
	ID     string `json:"id"`
	Tokens Tokens `json:"tokens"`
}

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

type VerifyCodeReq struct {
	ID          string `json:"id" gorm:"type:uuid;primary_key;"`
	PhoneNumber string `json:"phone" validate:"required,phone" gorm:"type:varchar(13);not null;unique;"`
	Code        string `json:"code" validate:"required,len=6" gorm:"-"`
}


func (req *VerifyCodeReq) Validate() error {
	if !utils.IsPhoneValid(req.PhoneNumber) {
		return errors.New("invalid phone number: must be in format +99XXXXXXXXXX")
	}
	if len(req.Code) != 6 {
		return errors.New("invalid code: must be 6 characters long")
	}
	return nil
}

type VerifyCodeRes struct {
	AccessToken string `json:"access_token"`
}

type LoginPostgres struct {
	Id          string `gorm:"primaryKey;column:id"`
	PhoneNumber string `gorm:"column:phone_number"`
}

type UpdateProfile struct {
	Firstname   *string    `json:"firstname" gorm:"column:firstname"`
	Surname     *string    `json:"surname" gorm:"column:surname"`
	Fathersname *string    `json:"fathersname" gorm:"column:fathersname"`
	Birthdate   *time.Time `json:"birthdate" gorm:"column:birthdate"`
	Gender      *string    `json:"gender" gorm:"column:gender"`
	CreatedBy   string    `json:"created_by" gorm:"column:created_by"`
	UpdatedBy   string    `json:"updated_by" gorm:"column:updated_by"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"column:updated_at"`
	CreatedAt   time.Time    `json:"created_at" gorm:"column:created_at"`
}
