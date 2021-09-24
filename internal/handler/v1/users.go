package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/errs"
	"github.com/denis-shcherbinin/spbpu-software-design-project/internal/service"
)

func (h *Handler) initUsersRoutes(api *echo.Group) {
	users := api.Group("/users")
	{
		auth := users.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
			auth.POST("/sign-out", h.signOut)
		}
	}
}

const minPasswordLength = 8

type signUpOpts struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func (opts *signUpOpts) Bind(c echo.Context) error {
	if err := c.Bind(opts); err != nil {
		return errors.New(fmt.Sprintf("can't bind: %v", err))
	}

	if len(opts.FirstName) == 0 {
		return errors.New("empty first_name")
	}

	if len(opts.Username) == 0 {
		return errors.New("empty username")
	}

	if len(opts.Password) < minPasswordLength {
		return errors.New(fmt.Sprintf("short password: length: %v; required: %v",
			len(opts.Password), minPasswordLength),
		)
	}

	return nil
}

func (h *Handler) signUp(c echo.Context) error {
	opts := signUpOpts{}
	if err := opts.Bind(c); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	user, err := h.services.Auth.SignUp(service.SignUpOpts{
		FirstName:  opts.FirstName,
		SecondName: opts.SecondName,
		Username:   opts.Username,
		Password:   opts.Password,
	})
	if err != nil {
		if err == errs.ErrUserAlreadyExists {
			return errorResponse(c, http.StatusBadRequest, err)
		}
		return errorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, echo.Map{"user": user})
}

func (h *Handler) signIn(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func (h *Handler) signOut(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}
