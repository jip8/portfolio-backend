package handlers

import (
	"net/http"
	"strconv"

	"github.com/jip/portfolio-backend/internal/api/experiences"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/labstack/echo/v4"
)

type ExperiencesHandler struct {
	useCase experiences.UseCase
}

func NewExperiencesHandler(useCase experiences.UseCase) *ExperiencesHandler {
	return &ExperiencesHandler{
		useCase: useCase,
	}
}

func (h *ExperiencesHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.ExperienceFlat
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		resp, err := h.useCase.Create(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
} 

func (h *ExperiencesHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.ExperienceFlat
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
		}

		req.Id = &id

		resp, err := h.useCase.Update(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func (h *ExperiencesHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
		}

		err = h.useCase.Delete(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func (h *ExperiencesHandler) ById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
		}

		resp, err := h.useCase.GetById(c.Request().Context(), id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func (h *ExperiencesHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.ListReq
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}
		
		resp, err := h.useCase.GetList(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}