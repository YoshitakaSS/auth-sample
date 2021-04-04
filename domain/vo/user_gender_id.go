package vo

import "errors"

// 不明（guest, corporation etc...
const unknown = 0

// 男性
const man = 1

// 女性
const female = 2

// UserGenderID はユーザーの性別IDを表す
type UserGenderID int

// NewUserGenderID UserGenderIDのポインタを返却
func NewUserGenderID(v int) (*UserGenderID, error) {
	l := []int{unknown, man, female}
	// 定義されていない値であれば
	if !inSlice(v, l) {
		return nil, errors.New("user gender value is ilegal")
	}

	r := UserGenderID(v)
	return &r, nil
}

func inSlice(s int, l []int) bool {
	for _, v := range l {
		if v == s {
			return true
		}
	}
	return false
}
