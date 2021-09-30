package v1

import (
	"encoding/base64"
	"fmt"

	"github.com/labstack/echo/v4"
)

func setBasicAuthHeader(c echo.Context, username, passwordHash string) {
	c.Response().Header().Set("Authorization", fmt.Sprintf("Basic %s", basicAuth(username, passwordHash)))
}

func basicAuth(username, passwordHash string) string {
	credentials := fmt.Sprintf("%s:%s", username, passwordHash)
	return base64.StdEncoding.EncodeToString([]byte(credentials))
}
