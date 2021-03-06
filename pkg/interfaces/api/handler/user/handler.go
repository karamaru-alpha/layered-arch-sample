package user

import (
	"fmt"
	"layered-arch-sample/pkg/constant"
	"layered-arch-sample/pkg/interfaces/api/dcontext"
	"layered-arch-sample/pkg/interfaces/api/myerror"
	uu "layered-arch-sample/pkg/usecase/user"
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// Handler UserにおけるHandlerのインターフェース
type Handler interface {
	HandleCreate(c echo.Context) error
	HandleGet(c echo.Context) error
	HandleUpdate(c echo.Context) error
}

type handler struct {
	useCase uu.UseCase
}

// NewHandler Userデータに関するHandlerを生成
func NewHandler(userUseCase uu.UseCase) Handler {
	return &handler{
		useCase: userUseCase,
	}
}

// HandleCreate ユーザを作成するHandler
func (uh handler) HandleCreate(c echo.Context) error {
	type (
		request struct {
			Name string `json:"name"`
		}
		response struct {
			Token string `json:"token"`
		}
	)

	requestBody := new(request)
	if err := c.Bind(requestBody); err != nil {
		return &myerror.InternalServerError{Err: err}
	}

	if len(requestBody.Name) < constant.MinNameLength || constant.MaxNameLength < len(requestBody.Name) {
		return &myerror.BadRequestError{
			Err:     errors.Errorf(`query "name" is invalid: name="%s"`, requestBody.Name),
			Message: fmt.Sprintf(`ユーザー名は2文字以上10文字以下に設定してください。(name: "%s")`, requestBody.Name),
		}
	}

	authToken, err := uh.useCase.Create(requestBody.Name)
	if err != nil {
		return &myerror.InternalServerError{Err: err}
	}

	return c.JSON(http.StatusOK, &response{Token: authToken})
}

// HandleGet ユーザー取得処理
func (uh handler) HandleGet(c echo.Context) error {
	type response struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	user := dcontext.GetUserFromContext(c)
	if user == nil {
		return &myerror.UnauthorizedError{Err: errors.New("user not found")}
	}

	return c.JSON(http.StatusOK, &response{
		ID:   user.ID,
		Name: user.Name,
	})
}

// HandleUpdate ユーザー更新処理
func (uh handler) HandleUpdate(c echo.Context) error {
	type (
		request struct {
			Name string `json:"name"`
		}
		response struct {
			Message string `json:"message"`
		}
	)

	requestBody := new(request)
	if err := c.Bind(requestBody); err != nil {
		return &myerror.InternalServerError{Err: err}
	}

	if len(requestBody.Name) < constant.MinNameLength || constant.MaxNameLength < len(requestBody.Name) {
		return &myerror.BadRequestError{
			Err:     errors.Errorf(`query "name" is invalid: name="%s"`, requestBody.Name),
			Message: fmt.Sprintf(`ユーザー名は2文字以上10文字以下に設定してください。(name: "%s")`, requestBody.Name),
		}
	}

	user := dcontext.GetUserFromContext(c)
	if user == nil {
		return &myerror.UnauthorizedError{Err: errors.New("user not found")}
	}

	if err := uh.useCase.UpdateName(user, requestBody.Name); err != nil {
		return &myerror.InternalServerError{Err: err}
	}

	return c.JSON(http.StatusOK, &response{
		Message: "Account successfully updated",
	})
}
