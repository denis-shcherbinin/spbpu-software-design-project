package handler

import (
	v1 "github.com/denis-shcherbinin/spbpu-software-design-project/internal/handler/v1"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

type InitOpts struct {
	Debug bool
}

func (h *Handler) Init(opts InitOpts) *echo.Echo {
	e := echo.New()

	// TODO: add basic auth middleware

	e.Use(middleware.CORS())

	e.Debug = opts.Debug

	e.GET("ping", Ping)

	h.initApi(e)

	return e
}

func (h *Handler) initApi(e *echo.Echo) {
	api := e.Group("/api")

	handlerV1 := v1.NewHandler(h.services)
	handlerV1.Init(api)
}
