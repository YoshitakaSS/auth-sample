package entity

import "github.com/YoshitakaSS/go_auth/domain/vo"

// User はユーザー情報を表す構造体
type User struct {
	Email     string
	userName  *vo.UserName
	birthDate *vo.UserBirthDate
	Password  string
	Gender    int
}

// NewUser はUserEntityを返します
func NewUser(u *vo.UserName, b *vo.UserBirthDate) (*User, error) {

	return &User{
		userName:  u,
		birthDate: b,
	}, nil
}
