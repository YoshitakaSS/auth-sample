package usecase

import (
	"github.com/YoshitakaSS/go_auth/domain/user/vo"
	"github.com/YoshitakaSS/go_auth/presenter/request"
	"github.com/YoshitakaSS/go_auth/usecase/query"
)

// UserFindUseCase はユースケース構造体
type UserFindUseCase struct {
	userRequest request.UserRegisterRequest
	userQs      query.UserQueryService
}

// UserFindUseCase はUserRegisterUseCaseのポインタを返却
func UserFindUseCase(userReq *request.UserRegisterRequest, userQs *query.UserQueryService) *UserFindUseCase {
	return &UserFindUseCase{userRequest: userReq, userQs: userQs}
}

// FindUser はユーザー情報を見つけるメソッド
func (u *UserFindUseCase) FindUser(string name) error {
	userName, _ := vo.NewUserName(name)

	dto := u.userQs.FindByUserName(userName)
	return dto
}
