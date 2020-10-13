package dcontext

import (
	um "layered-arch-sample/pkg/domain/model/user"

	"github.com/labstack/echo"
)

var userKey = "userKey"

// SetUser Contextへユーザを保存する
func SetUser(c echo.Context, user um.User) {
	c.Set(userKey, user)
}

// GetUserFromContext Contextからユーザを取得する
func GetUserFromContext(c echo.Context) *um.User {
	var user um.User
	if c.Get(userKey) == nil {
		return nil
	}

	user = c.Get(userKey).(um.User)
	return &user
}
