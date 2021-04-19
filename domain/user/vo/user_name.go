package vo

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

// UserName はユーザー名のデータ
type UserName string

// NewUserName は UserNameのポインタを返却
func NewUserName(v string) (*UserName, error) {
	re := regexp.MustCompile(`[A-Za-z]`)
	if !re.MatchString(v) {
		return nil, errors.New("user name is A-Z or a-z only")
	}

	// ユーザー名は5文字以上10文字以下
	if utf8.RuneCountInString(v) < 5 && utf8.RuneCountInString(v) > 10 {
		return nil, errors.New("user name is 5 to 10 characters")
	}

	un := UserName(v)
	return &un, nil
}
