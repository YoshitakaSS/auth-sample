package repository

import (
	"github.com/YoshitakaSS/go_auth/domain/user/repository"
	"github.com/go-gorp/gorp"
)

type UserRepository struct {
	dbMap *gorp.DbMap
}

func NewUserRepository(dbMap *gorp.DbMap) repository.UserRepository {
	return &UserRepository{dbMap: dbMap}
}
