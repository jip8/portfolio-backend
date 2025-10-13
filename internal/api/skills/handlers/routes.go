package handlers

import (
	"github.com/jip/portfolio-backend/internal/api/skills"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group, handler skills.Handlers, authMiddleware echo.MiddlewareFunc) {
	// Public routes
	e.GET("", handler.List())

	// Authenticated routes
	authGroup := e.Group("", authMiddleware)
	authGroup.PUT("", handler.Update())
}
