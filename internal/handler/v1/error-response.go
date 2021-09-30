package v1

import "github.com/labstack/echo/v4"

func errorResponse(c echo.Context, code int, err error) error {
	return c.JSON(code, echo.Map{
		"error": err.Error(),
	})
}
