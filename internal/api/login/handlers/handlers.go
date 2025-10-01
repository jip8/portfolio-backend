package handlers

import (
	"net/http"

	"github.com/jip/portfolio-backend/internal/api/login"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	useCase login.UseCase
}

func NewLoginHandler(useCase login.UseCase) *LoginHandler {
	return &LoginHandler{
		useCase: useCase,
	}
}

func (h *LoginHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.LoginRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		token, err := h.useCase.Login(c.Request().Context(), req, c.RealIP())
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"token": token})
	}
}
