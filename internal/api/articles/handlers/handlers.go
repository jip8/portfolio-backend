package handlers

import (
	"net/http"
	"strconv"
	"strings"

	portfolio "github.com/jip/portfolio-backend"
	"github.com/jip/portfolio-backend/internal/api/articles"
	"github.com/jip/portfolio-backend/internal/entity"
	"github.com/labstack/echo/v4"
	"github.com/jip/portfolio-backend/internal/api/attachments"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type ArticlesHandler struct {
	useCase articles.UseCase
	AttachmentsUC attachments.UseCase
}

func NewHandler(useCase articles.UseCase, AttachmentsUC attachments.UseCase) *ArticlesHandler {
	return &ArticlesHandler{
		useCase: useCase,
		AttachmentsUC: AttachmentsUC,
	}
}

func (h *ArticlesHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.ArticleFlat
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

func (h *ArticlesHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req entity.ArticleFlat
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

func (h *ArticlesHandler) Delete() echo.HandlerFunc {
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

func (h *ArticlesHandler) ById() echo.HandlerFunc {
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

func (h *ArticlesHandler) List() echo.HandlerFunc {
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

func (h *ArticlesHandler) InsertAttachment() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidIDFormat.Error()})
		}

		fileHeader, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "file is a required field"})
		}

		src, err := fileHeader.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to open file"})
		}
		defer src.Close()

		title := c.FormValue("title")
		if title == "" {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "title is a required field"})
		}
		description := c.FormValue("description")
		module := "articles"

		attachment := entity.AttachmentFlat{
			ParentId:    &id,
			Module:      &module,
			Title:       &title,
			Description: &description,
			FileObject: &entity.File{
				Name:        fileHeader.Filename,
				Size:        fileHeader.Size,
				ContentType: fileHeader.Header.Get("Content-Type"),
				Content:     src,
			},
		}

		attachments := []entity.AttachmentFlat{attachment}

		err = h.AttachmentsUC.Insert(c.Request().Context(), attachments)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return c.NoContent(http.StatusNoContent)
	}
}

func (h *ArticlesHandler) DeleteAttachment() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidIDFormat.Error()})
		}

		idsStr := c.QueryParam("ids")
		if idsStr == "" {
			return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "ids query param is required"})
		}

		idsSplit := strings.Split(idsStr, ",")
		var ids []int
		for _, idStr := range idsSplit {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return c.JSON(http.StatusBadRequest, ErrorResponse{Error: portfolio.ErrInvalidIDFormat.Error()})
			}
			ids = append(ids, id)
		}

		err = h.AttachmentsUC.Delete(c.Request().Context(), "articles", id, ids)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, ErrorResponse{Error: portfolio.ErrNotFound.Error()})
			}
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		}

		return c.NoContent(http.StatusNoContent)
	}
}