package v1

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
)

const (
	userIDCtx = "userID"
)

// noBody middleware checks that request has empty body
func noBody(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()

		if r.ContentLength > 0 {
			return echo.NewHTTPError(http.StatusRequestEntityTooLarge, "request body size must be zero")
		}

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("can't read request body: %v", err))
		}
		if len(body) > 0 {
			return echo.NewHTTPError(http.StatusRequestEntityTooLarge, "request body size must be zero")
		}

		r.Body = http.NoBody

		return next(c)
	}
}

func setBasicAuthHeader(c echo.Context, username, passwordHash string) {
	c.Response().Header().Set("Authorization", fmt.Sprintf("Basic %s", basicAuth(username, passwordHash)))
}

func basicAuth(username, passwordHash string) string {
	credentials := fmt.Sprintf("%s:%s", username, passwordHash)
	return base64.StdEncoding.EncodeToString([]byte(credentials))
}

func (h *Handler) basicAuthValidator(username, passwordHash string, c echo.Context) (bool, error) {
	userID, err := h.services.User.GetIDByCredentials(username, passwordHash)

	if err != nil {
		if err == errs.ErrUserNotFound {
			return false, nil
		}
		return false, err
	}

	c.Set(userIDCtx, userID)

	return true, nil
}

func (h *Handler) getUserID(c echo.Context) (int64, error) {
	idFromContext := c.Get(userIDCtx)

	id, ok := idFromContext.(int64)
	if !ok {
		return 0, fmt.Errorf("invalid user id format")
	}

	return id, nil
}
