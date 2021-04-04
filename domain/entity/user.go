package entity

import "github.com/YoshitakaSS/go_auth/domain/vo"

// User はユーザー情報を表す構造体
type User struct {
	Email     string
	userName  *vo.UserName
	birthDate *vo.UserBirthDate
	password  *vo.UserPassword
	Gender    int
}

// NewUser はUserEntityを返します
func NewUser(u *vo.UserName, b *vo.UserBirthDate, p *vo.UserPassword) (*User, error) {

	return &User{
		userName:  u,
		birthDate: b,
		password:  p,
	}, nil
}
