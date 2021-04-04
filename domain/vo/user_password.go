package vo

import (
	"errors"
	"unicode/utf8"
)

// UserPassword はユーザーパスワード
type UserPassword string

// NewUserPassword はUserPasswordのポインタを返却
func NewUserPassword(v string) (*UserPassword, error) {

	// パスワードは10文字以上、15文字以下のみ受け付ける
	if utf8.RuneCountInString(v) < 10 && utf8.RuneCountInString(v) > 15 {
		return nil, errors.New("password is 10 to 15 characters")
	}

	p := UserPassword(v)

	return &p, nil
}
