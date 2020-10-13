package user

// User Userを表すドメインモデル
type User struct {
	ID        string `json:"id"`
	AuthToken string `json:"auth_token"`
	Name      string `json:"name"`
}
