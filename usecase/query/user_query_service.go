package query

import (
	"github.com/YoshitakaSS/go_auth/domain/user/vo"
	"github.com/YoshitakaSS/go_auth/usecase/dto"
)

// UserQueryService はUser情報をのデータを取得する
type UserQueryService interface {
	FindByUserName(name *vo.UserName) (*dto.UserDTO, error)
}
