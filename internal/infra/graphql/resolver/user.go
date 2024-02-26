package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"fmt"

	"github.com/koki-algebra/go_server_sample/internal/entity"
	"github.com/koki-algebra/go_server_sample/internal/infra/graphql/generated/model"
)

// Save is the resolver for the Save field.
func (r *mutationResolver) Save(ctx context.Context, input model.SaveInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Save - Save"))
}

// GetByID is the resolver for the GetByID field.
func (r *queryResolver) GetByID(ctx context.Context, id string) (*model.User, error) {
	user, err := r.user.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func convertUser(user *entity.User) *model.User {
	if user == nil {
		return nil
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}
