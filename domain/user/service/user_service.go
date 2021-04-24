package service

import (
	"github.com/YoshitakaSS/go_auth/domain/user/vo"
	"github.com/YoshitakaSS/go_auth/infra/query"
)

// UserDomainService はドメインサービスの構造体
type UserDomainService struct {
	userQs *query.UserQueryService
}

// NewUserDomainService はUserDomainServiceの構造体を返す
func NewUserDomainService(userQs *query.UserQueryService) *UserDomainService {
	return &UserDomainService{userQs: userQs}
}

// IsExists はユーザーが存在するか判定する
func (s *UserDomainService) IsExists(name string) bool {
	userName, err := vo.NewUserName(name)
	if err != nil {
		return false
	}

	userDto, _ := s.userQs.FindByUserName(userName)

	if userDto == nil {
		return false
	}

	return true
}
