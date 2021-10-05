package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/service"
)

func (h *Handler) initTodoListsRoutes(api *echo.Group) {
	lists := api.Group("/lists", middleware.BasicAuth(h.basicAuthValidator))
	{
		lists.POST("/", h.createList)
		lists.GET("/", h.getAllLists)
		lists.GET("/:id", h.getList)
		lists.PUT("/:id", h.updateList)
		lists.DELETE("/:id", h.deleteList)
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
	if err := opts.Bind(c); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	err = h.services.List.Create(service.CreateListOpts{
		UserID:      userID,
		Title:       opts.Title,
		Description: opts.Description,
	})
	if err != nil {
		if err == errs.ErrListTitleAlreadyExists {
			return errorResponse(c, http.StatusBadRequest, fmt.Errorf("%v: %v", err, opts.Title))
		}
		return errorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"success": true,
	})
}

func (h *Handler) getAllLists(c echo.Context) error {
	return nil
}

func (h *Handler) getList(c echo.Context) error {
	return nil
}

func (h *Handler) updateList(c echo.Context) error {
	return nil
}

func (h *Handler) deleteList(c echo.Context) error {
	return nil
}
