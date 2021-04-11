package usecase

import (
	"github.com/YoshitakaSS/go_auth/domain/user/repository"
	"github.com/YoshitakaSS/go_auth/presenter/request"
)

// UserRegisterUseCase はユースケース構造体
type UserRegisterUseCase struct {
	userRequest request.UserRegisterRequest
	repo        repository.UserRepository
}

//
func NewUserRegisterUseCase(userReq *request.UserRegisterRequest, userRepo *repository.UserRepository) error {
	return &UserRegisterUseCase{userRequest: userReq, repo: userRepo}
}

func (u *UserRegisterUseCase) Register() error {

}
