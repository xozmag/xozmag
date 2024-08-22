package handlers

import (
	"context"
	"crypto/rand"
	"delivery/constants"
	"delivery/entities"
	"delivery/logger"
	htp "delivery/pkg/http"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	loginReq := entities.LoginReq{}
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		h.handleResponse(c, htp.BadRequest, err.Error())
		return
	}

	err = loginReq.Validate()
	if err != nil {
		h.handleResponse(c, htp.InvalidArgument, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), constants.ContextTimeoutDuration)
	defer cancel()

	resp, err := h.adminController.Login(
		ctx,
		loginReq,
	)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	h.handleResponse(c, htp.OK, resp)
}

func (h *Handler) SignUp(c *gin.Context) {
	req := entities.SendCodeReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, htp.BadRequest, err.Error())
		return
	}

	err = req.Validate()
	if err != nil {
		h.handleResponse(c, htp.InvalidArgument, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), constants.ContextTimeoutDuration)
	defer cancel()

	code, err := generateVerificationCode()
	if err != nil {
		h.log.Error("kod generatsiya qilishda xatolik", logger.Error(err))
		//return pkgerrors.NewError(http.StatusInternalServerError, "kod generatsiya qilishda xatolik")
	}

	redisKey := fmt.Sprintf("verification_code:%s", req.PhoneNumber)
	fmt.Println(redisKey)

	err = h.redis.Set(ctx, redisKey, code, 2*time.Minute).Err() // Kodni 2 minutga saqlaymiz
	if err != nil {
		h.log.Error("Redisda kodni saqlashda xatolik1", logger.Error(err))
		//return pkgerrors.NewError(http.StatusInternalServerError, "Redisda kodni saqlashda xatolik2")
	}

	h.handleResponse(c, htp.Created, "")
}

func (h *Handler) VerifyCode(c *gin.Context) {
	req := entities.VerifyCodeReq{}
	err := c.ShouldBindJSON(&req)

	if err != nil {
		h.handleResponse(c, htp.BadRequest, err.Error())
		return
	}

	err = req.Validate()
	if err != nil {
		h.handleResponse(c, htp.BadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), constants.ContextTimeoutDuration)
	defer cancel()

	isValid, err := h.adminController.VerifyCode(ctx, req)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	if !isValid {
		h.handleResponse(c, htp.Unauthorized, "Code is invalid or expired")
		return
	}

	h.handleResponse(c, htp.OK, "Code verified, you registered successfully!")
}

func generateVerificationCode() (string, error) {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	code := fmt.Sprintf("%06d", b[0]%(10*6))
	return code, nil
}
