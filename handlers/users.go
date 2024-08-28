package handlers

import (
	"context"
	"delivery/constants"
	"delivery/entities"
	"delivery/logger"
	htp "delivery/pkg/http"
	jwta "delivery/pkg/jwt"
	"delivery/pkg/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (u *Handler) UpdateProfile(c *gin.Context) {
	var req entities.UpdateProfile
	if err := c.ShouldBindJSON(&req); err != nil {
		u.handleResponse(c, htp.BadRequest, err.Error())
		return
	}

	userID, err := jwta.ExtractFromClaims("id", c.GetHeader("Auth"), []byte(u.cfg.JWTSecretKey))
	if err != nil {
		u.handleResponse(c, StatusFromError(err), err.Error())
		return
	}
	fmt.Println(userID, "//////////////")

	userIDStr, ok := userID.(string)
	if !ok {
		u.handleResponse(c, htp.InternalServerError, "Invalid user ID type")
		return
	}
	fmt.Println(userIDStr, "////////////////")

	req.UpdatedBy = userIDStr
	req.UpdatedAt = time.Now()

	if err := u.adminController.UpdateUserProfile(c.Request.Context(), userIDStr, req); err != nil {
		u.handleResponse(c, htp.InternalServerError, err.Error())
		return
	}

	u.handleResponse(c, htp.OK, "Profile updated successfully!")
}

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

func (h *Handler) SendCode(c *gin.Context) {
	var req entities.SendCodeReq
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

	smscode, err := utils.GenerateVerificationCode()
	if err != nil {
		h.log.Error("kod generatsiya qilishda xatolik", logger.Error(err))
	}

	err = h.redis.Set(ctx, req.PhoneNumber, smscode, 20*time.Minute).Err()
	if err != nil {
		h.log.Error("Redisda kodni saqlashda xatolik1", logger.Error(err))
	}

	h.handleResponse(c, htp.Created, "")
}

func (h *Handler) SignUp(c *gin.Context) {
	var req entities.VerifyCodeReq
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

	isValid, err := h.adminController.SignUp(ctx, req)
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
