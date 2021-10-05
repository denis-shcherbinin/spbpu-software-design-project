package v1

import (
	"encoding/base64"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
)

const (
	userIDCtx = "userID"
)

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
