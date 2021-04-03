package vo

import (
	"errors"
	"regexp"
)

// UserBirthDate はユーザーの生年月日
type UserBirthDate string

// NewUserBirthDate はUserBirthDateのポインタを返却
func NewUserBirthDate(v string) (*UserBirthDate, error) {
	re := regexp.MustCompile(`[0-9]{4}/(0[1-9]|1[0-2])/(0[1-9]|[12][0-9]|3[01])`)

	if !re.MatchString(v) {
		return nil, errors.New("user birth date yyyy-mm-dd format")
	}

	ub := UserBirthDate(v)
	return &ub, nil
}
