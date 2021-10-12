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
	users := api.Group("/user")
	{
		auth := users.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}
	}
}

const (
	minUsernameLength = 4
	maxUsernameLength = 32
	minPasswordLength = 8
	maxPasswordLength = 32
)

var (
	ErrInvalidUsernameLength = errors.New("invalid username length")
	ErrInvalidPasswordLength = errors.New("invalid password length")
)

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

	if !niceUsername(opts.Username) {
		return ErrInvalidUsernameLength
	}

	if !nicePassword(opts.Password) {
		return ErrInvalidPasswordLength
	}

	return nil
}

func (h *Handler) signUp(c echo.Context) error {
	opts := signUpOpts{}
	if err := opts.Bind(c); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	err := h.services.Auth.SignUp(service.SignUpOpts{
		FirstName:  opts.FirstName,
		SecondName: opts.SecondName,
		Username:   opts.Username,
		Password:   opts.Password,
	})
	if err != nil {
		if err == errs.ErrUserAlreadyExists {
			return errorResponse(c, http.StatusConflict, fmt.Errorf("AuthService.SignUp: %v", err))
		}
		return errorResponse(c, http.StatusInternalServerError, fmt.Errorf("AuthService.SignUp: %v", err))
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"success": true},
	)
}

type SignInOpts struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (opts *SignInOpts) Bind(c echo.Context) error {
	if err := c.Bind(opts); err != nil {
		return errors.New(fmt.Sprintf("can't bind: %v", err))
	}

	if !niceUsername(opts.Username) {
		return ErrInvalidUsernameLength
	}

	if !nicePassword(opts.Password) {
		return ErrInvalidPasswordLength
	}

	return nil
}

func (h *Handler) signIn(c echo.Context) error {
	opts := &SignInOpts{}
	if err := opts.Bind(c); err != nil {
		return errorResponse(c, http.StatusBadRequest, err)
	}

	username, passwordHash, err := h.services.Auth.SignIn(service.SignInOpts{
		Username: opts.Username,
		Password: opts.Password,
	})
	if err != nil {
		if err == errs.ErrUserNotFound {
			return errorResponse(c, http.StatusForbidden, fmt.Errorf("AuthService.SignIn: %v", err))
		}
		return errorResponse(c, http.StatusInternalServerError, fmt.Errorf("AuthService.SignIn: %v", err))
	}

	setBasicAuthHeader(c, username, passwordHash)

	return c.JSON(http.StatusOK, echo.Map{
		"success": true,
	})
}

func nicePassword(password string) bool {
	length := len(password)
	return minPasswordLength <= length && length <= maxPasswordLength
}

func niceUsername(username string) bool {
	length := len(username)
	return minUsernameLength <= length && length <= maxUsernameLength
}
