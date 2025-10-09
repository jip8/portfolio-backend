package handlers

import (
	"github.com/jip/portfolio-backend/internal/api/projects"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group, handler projects.Handlers, authMiddleware echo.MiddlewareFunc) {
	// Public routes
	e.GET("/:id", handler.ById())
	e.GET("", handler.List())

	// Authenticated routes
	authGroup := e.Group("", authMiddleware)
	authGroup.POST("", handler.Create())
	authGroup.PUT("/:id", handler.Update())
	authGroup.DELETE("/:id", handler.Delete())
}
