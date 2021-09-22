package v1

import (
	"github.com/labstack/echo/v4"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *echo.Group) {
	v1 := api.Group("/v1")
	h.initUsersRoutes(v1)
}
