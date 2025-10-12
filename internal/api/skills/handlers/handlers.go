package handlers

import (
	"net/http"
	"strings"

	portfolio "github.com/jip/portfolio-backend"
	"github.com/jip/portfolio-backend/internal/api/skills"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SkillsHandler struct {
	useCase skills.UseCase
}

func NewHandler(useCase skills.UseCase) *SkillsHandler {
	return &SkillsHandler{
		useCase: useCase,
	}
}

func (h *SkillsHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.SkillArray
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidRequestBody.Error()})
		}

		err := h.useCase.Upsert(c.Request().Context(), nil, nil, req)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, ErrorResponse{Error: portfolio.ErrNotFound.Error()})
			}
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		resp, err := h.useCase.GetList(c.Request().Context(), nil, nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func (h *SkillsHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {

		resp, err := h.useCase.GetList(c.Request().Context(), nil, nil)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}