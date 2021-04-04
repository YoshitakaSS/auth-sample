package vo

import (
	"errors"
	"unicode/utf8"
)

// UserName はユーザー名のデータ
type UserName string

// NewUserName は UserNameのポインタを返却
func NewUserName(v string) (*UserName, error) {
	// ユーザー名は5文字以上10文字以下
	if utf8.RuneCountInString(v) < 5 && utf8.RuneCountInString(v) > 10 {
		return nil, errors.New("user name is 5 to 10 characters")
	}
	un := UserName(v)
	return &un, nil
}
