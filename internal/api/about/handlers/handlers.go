package handlers

import (
	"net/http"
	"strings"

	portfolio "github.com/jip/portfolio-backend"
	"github.com/jip/portfolio-backend/internal/api/about"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type AboutHandler struct {
	useCase about.UseCase
}

func NewHandler(useCase about.UseCase) *AboutHandler {
	return &AboutHandler{
		useCase: useCase,
	}
}

func (h *AboutHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.About
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidRequestBody.Error()})
		}

		resp, err := h.useCase.Update(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func (h *AboutHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		resp, err := h.useCase.Get(c.Request().Context())
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, ErrorResponse{Error: portfolio.ErrNotFound.Error()})
			}
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}
