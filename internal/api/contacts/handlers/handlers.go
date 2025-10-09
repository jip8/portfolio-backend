package handlers

import (
	"net/http"
	"strconv"
	"strings"

	portfolio "github.com/jip/portfolio-backend"
	"github.com/jip/portfolio-backend/internal/api/contacts"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type ContactsHandler struct {
	useCase contacts.UseCase
}

func NewHandler(useCase contacts.UseCase) *ContactsHandler {
	return &ContactsHandler{
		useCase: useCase,
	}
}

func (h *ContactsHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.Contact
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

func (h *ContactsHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.Contact
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

func (h *ContactsHandler) Delete() echo.HandlerFunc {
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

func (h *ContactsHandler) ById() echo.HandlerFunc {
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

func (h *ContactsHandler) List() echo.HandlerFunc {
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
