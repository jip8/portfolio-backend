package handlers

import (
	"github.com/jip/portfolio-backend/internal/api/experiences"
	"github.com/labstack/echo/v4"
)

func ExperiencesRoutes(e *echo.Group, handler experiences.Handlers) {
	e.POST("", handler.Create())
	e.PUT("/:id", handler.Update())
	e.DELETE("/:id", handler.Delete())

	e.GET("/:id", handler.ById())

	e.GET("", handler.List())

	// e.GET("/transitions", handler.GetTransitions())
}