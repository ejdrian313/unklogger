package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"unklogger.com/src/application/services"
)

type AuthController struct {
	service *services.AuthServices
}

func NewAuthController(e *echo.Echo, service *services.AuthServices) {
	controller := &AuthController{
		service: service,
	}

	e.POST("/login", controller.Login)
	e.POST("/register", controller.Register)
}

func (ac *AuthController) Login(c echo.Context) error {
	loginRequest := &LoginRequest{}
	if err := c.Bind(loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}
	token, err := ac.service.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid username or password",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (ac *AuthController) Register(c echo.Context) error {
	registerRequest := &RegisterRequest{}
	if err := c.Bind(registerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}
	token, err := ac.service.Register(registerRequest.Username, registerRequest.Password)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "Con not register user with that username",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
