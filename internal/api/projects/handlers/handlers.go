package handlers

import (
	"net/http"
	"strconv"
	"strings"

	portfolio "github.com/jip/portfolio-backend"
	"github.com/jip/portfolio-backend/internal/api/projects"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type ProjectsHandler struct {
	useCase projects.UseCase
}

func NewHandler(useCase projects.UseCase) *ProjectsHandler {
	return &ProjectsHandler{
		useCase: useCase,
	}
}

func (h *ProjectsHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.ProjectFlat
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidRequestBody.Error()})
		}

		resp, err := h.useCase.Create(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func (h *ProjectsHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.ProjectFlat
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidRequestBody.Error()})
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidIDFormat.Error()})
		}

		req.Id = &id

		resp, err := h.useCase.Update(c.Request().Context(), req)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, ErrorResponse{Error: portfolio.ErrNotFound.Error()})
			}
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func (h *ProjectsHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidIDFormat.Error()})
		}

		err = h.useCase.Delete(c.Request().Context(), id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, ErrorResponse{Error: portfolio.ErrNotFound.Error()})
			}
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func (h *ProjectsHandler) ById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidIDFormat.Error()})
		}

		resp, err := h.useCase.GetById(c.Request().Context(), id)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, ErrorResponse{Error: portfolio.ErrNotFound.Error()})
			}
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}

func (h *ProjectsHandler) List() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.ListReq

		limitStr := c.QueryParam("limit")
		if limitStr != "" {
			limit, err := strconv.Atoi(limitStr)
			if err != nil {
				return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidLimitFormat.Error()})
			}
			req.Limit = limit
		}

		offsetStr := c.QueryParam("offset")
		if offsetStr != "" {
			offset, err := strconv.Atoi(offsetStr)
			if err != nil {
				return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidOffsetFormat.Error()})
			}
			req.Offset = offset
		}

		req.Process()

		resp, err := h.useCase.GetList(c.Request().Context(), req)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}
