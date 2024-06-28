package repository

type User struct {
	UserID         uint   `json:"user_id" gorm:"primary_key"` // 用户ID
	Username       string `json:"username" gorm:"unique"`     // 用户名
	NickName       string `json:"nickname"`                   // 昵称
	PasswordDigest string `json:"password_digest"`            // 加密后的密码
}
