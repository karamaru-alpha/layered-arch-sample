package user

import (
	"database/sql"
	ur "layered-arch-sample/pkg/domain/repository/db/user"
)

type repositoryImpl struct {
	db *sql.DB
}

// NewRepositoryImpl Userに関するDB更新処理を生成
func NewRepositoryImpl(db *sql.DB) ur.Repository {
	return &repositoryImpl{
		db,
	}
}

// Create ユーザ登録処理
func (uri repositoryImpl) Create(ID, authToken, name string) error {
	stmt, err := uri.db.Prepare("INSERT INTO users (id, auth_token, name) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(ID, authToken, name)
	return err
}
