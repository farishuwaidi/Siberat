package handler

import (
	"Siberat/dto"
	errorhandler "Siberat/errorHandler"
	"Siberat/helper"
	"Siberat/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}
	if err := h.service.Register(&register); err != nil {
		errorhandler.HandlerError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message: "RegisterSuccessfuly",
	})

	c.JSON(http.StatusCreated, res)
}

func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		errorhandler.HandlerError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.Login(&login)
	if err !=  nil {
		errorhandler.HandlerError(c,err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message: "Login success",
		Data: result,
	})

	c.JSON(http.StatusOK, res)
}