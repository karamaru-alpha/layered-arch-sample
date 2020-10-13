package server

import (
	"fmt"
	"layered-arch-sample/pkg/infrastructure/mysql"
	ur "layered-arch-sample/pkg/infrastructure/mysql/repositoryimpl/user"
	uh "layered-arch-sample/pkg/interfaces/api/handler/user"
	"layered-arch-sample/pkg/interfaces/api/myerror"
	uu "layered-arch-sample/pkg/usecase/user"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Serve サーバー起動処理
func Serve(addr string) {

	// 依存性の注入
	userRepoImpl := ur.NewRepositoryImpl(mysql.Conn)
	userUseCase := uu.NewUseCase(userRepoImpl)
	userHandler := uh.NewHandler(userUseCase)

	echo.NotFoundHandler = func(c echo.Context) error {
		return &myerror.NotFoundError{Err: fmt.Errorf("URL is invalid: (url=%s)", c.Request().URL)}
	}
	e := echo.New()
	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"Content-Type", "Accept", "Origin", "x-token"},
		}),
		middleware.Logger(),
		middleware.Recover(),
	)
	e.HTTPErrorHandler = errorHandler

	e.POST("/signup", userHandler.HandleCreate)

	log.Println("Server running...")
	if err := e.Start(addr); err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

func errorHandler(err error, c echo.Context) {
	type response struct {
		Message string `json:"message"`
	}
	var (
		code    int
		msg     string
		errInfo error
	)

	switch e := err.(type) {
	case *myerror.BadRequestError:
		code = http.StatusBadRequest
		msg = e.Message
		errInfo = e.Err
	case *myerror.UnauthorizedError:
		code = http.StatusUnauthorized
		msg = "401 認証エラー"
		errInfo = e.Err
	case *myerror.NotFoundError:
		code = http.StatusNotFound
		msg = "404 not found"
		errInfo = e.Err
	case *myerror.InternalServerError:
		code = http.StatusInternalServerError
		msg = "内部的なエラーが発生しました。リトライしてみてください。"
		errInfo = e.Err
	default:
		code = http.StatusInternalServerError
		msg = "エラーが発生しました。リトライしてみてください。"
		errInfo = err
	}

	// Todo: 認証後はUserIDも表示
	log.Printf(`access:"%s", errorCode:%d, errorMessage:"%s", error="%+v"`, c.Request().URL, code, msg, errInfo)

	if !c.Response().Committed {
		if err := c.JSON(code, &response{
			Message: msg,
		}); err != nil {
			log.Print("errorResponseのJson変換エラー")
		}
	}
}
