package handlers

import (
	"github.com/jip/portfolio-backend/internal/api/login"
	"github.com/labstack/echo/v4"
)

func LoginRoutes(e *echo.Group, handler login.Handlers) {
	e.POST("", handler.Login())
}