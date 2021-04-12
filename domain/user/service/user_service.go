package service

import (
	"errors"

	"github.com/YoshitakaSS/go_auth/domain/user/repository"
	"github.com/YoshitakaSS/go_auth/domain/user/vo"
)

// UserDomainService はドメインサービスの構造体
type UserDomainService struct {
	repo repository.UserRepository
}

// NewUserDomainService はUserDomainServiceの構造体を返す
func NewUserDomainService(repo *repository.UserRepository) error {
	return &UserDomainService{repo: repo}
}

// IsExists はユーザーが存在するか判定する
func (s *UserDomainService) IsExists(name *vo.UserName) error {
	_, ok := s.repo.Find(name)

	if !ok {
		return errors.New("not find this user")
	}

	return ok
}
