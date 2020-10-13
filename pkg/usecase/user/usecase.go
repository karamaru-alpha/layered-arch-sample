package user

import (
	um "layered-arch-sample/pkg/domain/model/user"
	ur "layered-arch-sample/pkg/domain/repository/db/user"

	"github.com/google/uuid"
)

// UseCase Userにおけるユースケースのインターフェース
type UseCase interface {
	Create(name string) (authToken string, err error)
	SelectByAuthToken(authToken string) (user *um.User, err error)
}

type useCase struct {
	repository ur.Repository
}

// NewUseCase Userデータに関するユースケースを生成
func NewUseCase(userRepo ur.Repository) UseCase {
	return &useCase{
		repository: userRepo,
	}
}

// CreateUser Userを新規作成するためのユースケース
func (uu useCase) Create(name string) (string, error) {
	userID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	authToken, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	if err := uu.repository.Create(userID.String(), authToken.String(), name); err != nil {
		return "", err
	}

	return authToken.String(), nil
}

// SelectByAuthToken Userをトークンから取得するためのユースケース
func (uu useCase) SelectByAuthToken(authToken string) (*um.User, error) {
	return uu.repository.SelectByAuthToken(authToken)
}
