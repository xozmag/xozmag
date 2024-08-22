package entities

import (
	"errors"
	"delivery/pkg/utils"
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

type SignupReq struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone" validate:"required,phone"`
}

func (req *SignupReq) Validate() error {
	if !utils.IsPhoneValid(req.PhoneNumber) {
		return errors.New("invalid phone number")
	}
	return nil
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
	PhoneNumber string `json:"phone" validate:"required,phone"`
	Code        string `json:"code" validate:"required,len=6"`
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
