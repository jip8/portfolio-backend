package handlers

import (
	"github.com/jip/portfolio-backend/internal/api/experiences"
	"github.com/labstack/echo/v4"
)

func ExperiencesRoutes(e *echo.Group, handler experiences.Handlers, authMiddleware echo.MiddlewareFunc) {
	// Public routes
	e.GET("/:id", handler.ById())
	e.GET("", handler.List())

	// Authenticated routes
	authGroup := e.Group("", authMiddleware)
	authGroup.POST("", handler.Create())
	authGroup.PUT("/:id", handler.Update())
	authGroup.DELETE("/:id", handler.Delete())
}