package vo

import (
	"errors"
	"regexp"
)

// UserEmail はユーザーのメールアドレス
type UserEmail string

// NewUserEmail はUserEmailのポインタを返却
func NewUserEmail(v string) (*UserEmail, error) {
	re := regexp.MustCompile(`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`)

	if !re.MatchString(v) {
		return nil, errors.New("user email formart is illegal")
	}

	e := UserEmail(v)
	return &e, nil
}
