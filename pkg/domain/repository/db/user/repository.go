package user

import (
	"layered-arch-sample/pkg/domain/model/user"
)

// Repository データ永続化のために抽象化したUserデータ更新周りの処理
type Repository interface {
	Create(ID, authToken, name string) error
	SelectByAuthToken(authToken string) (*user.User, error)
}
