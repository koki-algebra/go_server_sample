package resolver

import (
	"github.com/koki-algebra/go_server_sample/internal/entity"
	"github.com/koki-algebra/go_server_sample/internal/infra/graphql/generated/model"
)

func convertUser(user *entity.User) *model.User {
	if user == nil {
		return nil
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}
