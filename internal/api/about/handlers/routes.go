package handlers

import (
	"github.com/jip/portfolio-backend/internal/api/about"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group, handler about.Handlers, authMiddleware echo.MiddlewareFunc) {
	// Public routes
	e.GET("", handler.Get())

	// Authenticated routes
	authGroup := e.Group("", authMiddleware)
	authGroup.PUT("", handler.Update())
}
