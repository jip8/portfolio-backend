package handlers

import (
	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group, handler articles.Handlers, authMiddleware echo.MiddlewareFunc) {
	// Public routes
	e.GET("/:id", handler.ById())
	e.GET("", handler.List())

	// Authenticated routes
	authGroup := e.Group("", authMiddleware)
	authGroup.POST("", handler.Create())
	authGroup.PUT("/:id", handler.Update())
	authGroup.DELETE("/:id", handler.Delete())
	authGroup.POST("/:id/attachments", handler.InsertAttachment())
	authGroup.DELETE("/:id/attachments", handler.DeleteAttachment())
}
