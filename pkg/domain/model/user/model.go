package user

// User Userを表すドメインモデル
type User struct {
	ID        string
	AuthToken string
	Name      string
}
