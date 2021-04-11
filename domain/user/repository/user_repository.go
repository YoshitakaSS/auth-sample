package repository

import "github.com/YoshitakaSS/go_auth/domain/user/entity"

// UserRepository はUser関連のデータ永続化担当する
type UserRepository interface {
	Register(user *entity.User) error
}
