package query

import (
	"fmt"

	"github.com/YoshitakaSS/go_auth/domain/user/vo"
	"github.com/YoshitakaSS/go_auth/usecase/dto"
	"github.com/YoshitakaSS/go_auth/usecase/query"
	"github.com/go-gorp/gorp"
)

// UserQueryService はユーザー参照
type UserQueryService struct {
	dbMap *gorp.DbMap
}

// NewUserQueryService はUserQueryServiceの構造体を返却
func NewUserQueryService(dbMap *gorp.DbMap) query.UserQueryService {
	return &UserQueryService{dbMap: dbMap}
}

func (qs *UserQueryService) FindByUserName(name *vo.UserName) (*dto.UserDTO, error) {
	var dto dto.UserDTO

	query := `
		SELECT user_id, user_name, email, birth_date
		FROM users
		WHERE username = ?
	`

	if err := qs.dbMap.SelectOne(&dto, query, name); err != nil {
		return nil, fmt.Errorf("not found select user %s", name)
	}

	return &dto, nil
}
