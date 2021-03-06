package v1

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/service"
)

func (h *Handler) initTodoListsRoutes(api *echo.Group) {
	lists := api.Group("/lists", middleware.BasicAuth(h.basicAuthValidator))
	{
		lists.POST("/", h.createList)
		lists.GET("/", h.getAllLists, noBody)
		lists.GET("/:id", h.getList, noBody)
		lists.PUT("/:id", h.updateList)
		lists.DELETE("/:id", h.deleteList, noBody)

		items := lists.Group("/:id/items")
		{
			items.POST("/", h.createItem)
			items.GET("/", h.getAllItems, noBody)
		}
	}
}

type createListOpts struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (opts *createListOpts) Bind(c echo.Context) error {
	if err := c.Bind(opts); err != nil {
		return fmt.Errorf("can't bind: %v", err)
	}

	if len(opts.Title) == 0 {
		return errors.New("empty todo-list title")
	}

	return nil
}

func (h *Handler) createList(c echo.Context) error {
	userID, err := h.getUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	opts := &createListOpts{}
	if err = opts.Bind(c); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	err = h.services.List.Create(userID, service.CreateListOpts{
		Title:       opts.Title,
		Description: opts.Description,
	})
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, fmt.Errorf("ListService.Create: %v", err))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"success": true,
	})
}

func (h *Handler) getAllLists(c echo.Context) error {
	userID, err := h.getUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	lists, err := h.services.List.GetAll(userID)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, fmt.Errorf("ListService.GetAll: %v", err))
	}

	return c.JSON(http.StatusOK, lists)
}

func (h *Handler) getList(c echo.Context) error {
	listID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	userID, err := h.getUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	list, err := h.services.List.GetByID(userID, listID)
	if err != nil {
		if err == errs.ErrListNotFound {
			return errorResponse(c, http.StatusBadRequest, fmt.Errorf("ListService.GetByID: %v: %v", err, listID))
		}
		return errorResponse(c, http.StatusInternalServerError, fmt.Errorf("ListService.GetByID: %v", err))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"list": list,
	})
}

type updateListOpts struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (opts *updateListOpts) Bind(c echo.Context) error {
	if err := c.Bind(opts); err != nil {
		return fmt.Errorf("can't bind: %v", err)
	}

	if opts.Title == nil && opts.Description == nil {
		return errors.New("empty list update input: all parameters = null")
	}

	if opts.Title != nil && opts.Description != nil {
		if len(*opts.Title) == 0 && len(*opts.Description) == 0 {
			return errors.New("empty list update input")
		}
	}

	return nil
}

func (h *Handler) updateList(c echo.Context) error {
	listID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	userID, err := h.getUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	opts := &updateListOpts{}
	if err = opts.Bind(c); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	err = h.services.List.Update(userID, listID, service.UpdateListOpts{
		Title:       opts.Title,
		Description: opts.Description,
	})
	if err != nil {
		if err == errs.ErrListNotFound {
			return errorResponse(c, http.StatusBadRequest, fmt.Errorf("ListService.Update: %v: %v", err, listID))
		}
		return errorResponse(c, http.StatusInternalServerError, fmt.Errorf("ListService.Update: %v", err))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
	})
}

func (h *Handler) deleteList(c echo.Context) error {
	listID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	userID, err := h.getUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	err = h.services.List.DeleteByID(userID, listID)
	if err != nil {
		if err == errs.ErrListNotFound {
			return errorResponse(c, http.StatusBadRequest,
				fmt.Errorf("ListService.DeleteByID: %v: %v", err, listID))
		}
		return errorResponse(c, http.StatusInternalServerError, fmt.Errorf("ListService.DeleteByID: %v", err))
	}

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
	})
}

type createItemOpts struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (opts *createItemOpts) Bind(c echo.Context) error {
	if err := c.Bind(opts); err != nil {
		return fmt.Errorf("can't bind: %v", err)
	}

	if len(opts.Title) == 0 {
		return errors.New("empty todo-item title")
	}

	return nil
}

func (h *Handler) createItem(c echo.Context) error {
	listID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	userID, err := h.getUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	opts := &createItemOpts{}
	if err := opts.Bind(c); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	err = h.services.Item.Create(userID, listID, service.CreateItemOpts{
		Title:       opts.Title,
		Description: opts.Description,
	})

	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, fmt.Errorf("ItemService.Create: %v", err))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"success": true,
	})
}

func (h *Handler) getAllItems(c echo.Context) error {
	listID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	userID, err := h.getUserID(c)
	if err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	items, err := h.services.Item.GetAll(userID, listID)
	if err != nil {
		return errorResponse(c, http.StatusInternalServerError, fmt.Errorf("ItemService.GetAll: %v", err))
	}

	return c.JSON(http.StatusOK, items)
}
