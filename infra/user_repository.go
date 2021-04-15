package infra

import "github.com/YoshitakaSS/go_auth/domain/user/repository"

type UserRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &UserRepository
}

func (r *UserRepository) FindUser() {

}
