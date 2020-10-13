package user

// Repository データ永続化のために抽象化したUserデータ更新周りの処理
type Repository interface {
	Create(ID, authToken, name string) error
}
