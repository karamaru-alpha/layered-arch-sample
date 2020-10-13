package user

import (
	"database/sql"
	um "layered-arch-sample/pkg/domain/model/user"
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

// SelectByAuthToken auth_tokenを条件にUserを取得する
func (uri repositoryImpl) SelectByAuthToken(authToken string) (*um.User, error) {
	row := uri.db.QueryRow("SELECT * FROM users WHERE auth_token = ?", authToken)
	return convertToUser(row)
}

func convertToUser(row *sql.Row) (*um.User, error) {
	user := um.User{}
	if err := row.Scan(&user.ID, &user.AuthToken, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
