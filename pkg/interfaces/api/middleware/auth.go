package middleware

import (
	"layered-arch-sample/pkg/interfaces/api/dcontext"
	"layered-arch-sample/pkg/interfaces/api/myerror"
	"layered-arch-sample/pkg/usecase/user"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// Middleware middlewareのインターフェース
type Middleware interface {
	Authenticate(echo.HandlerFunc) echo.HandlerFunc
}

type middleware struct {
	useCase user.UseCase
}

// NewMiddleware userUseCaseと疎通
func NewMiddleware(uu user.UseCase) Middleware {
	return &middleware{
		useCase: uu,
	}
}

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func (m middleware) Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Request().Header.Get("x-token")
		if token == "" {
			return &myerror.UnauthorizedError{Err: errors.New("x-token is empty")}
		}

		user, err := m.useCase.SelectByAuthToken(token)
		if err != nil {
			return &myerror.InternalServerError{Err: err}
		}
		if user == nil {
			return &myerror.UnauthorizedError{Err: errors.Errorf(`user is not found: token="%s"`, token)}
		}

		dcontext.SetUser(c, *user)

		return next(c)
	}
}
