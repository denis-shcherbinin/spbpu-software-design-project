package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/service"
)

func (h *Handler) initTodoItemsRoutes(api *echo.Group) {
	items := api.Group("/items", middleware.BasicAuth(h.basicAuthValidator))
	{
		items.GET("/:id", h.getItem, noBody)
		items.PUT("/:id", h.updateItem)
		items.DELETE("/:id", h.deleteItem, noBody)
	}
}

func (h *Handler) getItem(c echo.Context) error {
	itemID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	userID, err := h.getUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	item, err := h.services.Item.GetByID(userID, itemID)
	if err != nil {
		// TODO: [todo-lists] not found error handle
		return errorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"item": item,
	})
}

type updateItemOpts struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (opts *updateItemOpts) Bind(c echo.Context) error {
	if err := c.Bind(opts); err != nil {
		return fmt.Errorf("can't bind: %v", err)
	}

	if opts.Title == nil && opts.Description == nil && opts.Done == nil {
		return errors.New("empty item update input: all parameters = null")
	}

	if opts.Title != nil && opts.Description != nil {
		if len(*opts.Title) == 0 && len(*opts.Description) == 0 && opts.Done == nil {
			return errors.New("empty item update input")
		}
	}

	return nil
}

func (h *Handler) updateItem(c echo.Context) error {
	itemID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	userID, err := h.getUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	opts := &updateItemOpts{}
	if err := opts.Bind(c); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	err = h.services.Item.Update(userID, itemID, service.UpdateItemOpts{
		Title:       opts.Title,
		Description: opts.Description,
		Done:        opts.Done,
	})

	if err != nil {
		// TODO: [todo-lists] not found error handle
		return errorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
	})
}

func (h *Handler) deleteItem(c echo.Context) error {
	itemID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	userID, err := h.getUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	err = h.services.Item.DeleteByID(userID, itemID)
	if err != nil {
		// TODO: [todo-lists] not found error handle
		return errorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
	})
}
