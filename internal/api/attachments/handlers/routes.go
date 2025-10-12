package handlers

import (
	"github.com/jip/portfolio-backend/internal/api/attachments"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Group, handler attachments.Handlers, authMiddleware echo.MiddlewareFunc) {
	// Public routes
	e.GET("/:uuid", handler.Get())
}
