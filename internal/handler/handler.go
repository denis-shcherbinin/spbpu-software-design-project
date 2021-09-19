package handler

import "github.com/labstack/echo/v4"

func InitRoutes(e *echo.Echo) error {
	e.GET("ping", Ping)

	return nil
}
