package entity

import "github.com/YoshitakaSS/go_auth/domain/vo"

// User はユーザー情報を表す構造体
type User struct {
	email    *vo.UserEmail
	name     *vo.UserName
	birth    *vo.UserBirthDate
	password *vo.UserPassword
	gender   *vo.UserGenderID
}

// NewUser はUserEntityを返します
func NewUser(e *vo.UserEmail, u *vo.UserName, b *vo.UserBirthDate, p *vo.UserPassword, g *vo.UserGenderID) *User {
	return &User{
		email:    e,
		name:     u,
		birth:    b,
		password: p,
		gender:   g,
	}
}
