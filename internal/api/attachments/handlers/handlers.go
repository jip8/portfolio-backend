package handlers

import (
	"net/http"
	"strings"

	portfolio "github.com/jip/portfolio-backend"
	"github.com/jip/portfolio-backend/internal/api/attachments"
	"github.com/labstack/echo/v4"
	"io"
	"fmt"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type AttachmentsHandler struct {
	useCase attachments.UseCase
}

func NewHandler(useCase attachments.UseCase) *AttachmentsHandler {
	return &AttachmentsHandler{
		useCase: useCase,
	}
}


func (h *AttachmentsHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		uuid := c.Param("uuid")

		resp, err := h.useCase.Get(c.Request().Context(), uuid)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, ErrorResponse{Error: portfolio.ErrNotFound.Error()})
			}
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		c.Response().Header().Set(echo.HeaderContentType, resp.ContentType)
		c.Response().Header().Set(echo.HeaderContentLength, fmt.Sprintf("%d", resp.Size))
		c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("inline; filename=\"%s\"", resp.Name))

		_, err = io.Copy(c.Response().Writer, resp.Content)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return nil
	}
}
