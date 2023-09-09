package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/koki-algebra/go_server_sample/internal/entity"
	"github.com/koki-algebra/go_server_sample/internal/infra/http/generated"
	"github.com/koki-algebra/go_server_sample/internal/usecase"
)

type User struct {
	usecase *usecase.User
}

func NewUser(usecase *usecase.User) *User {
	return &User{
		usecase: usecase,
	}
}

func (h *User) GetUserByID(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()
	user, err := h.usecase.GetByID(ctx, id)
	if err != nil {
		err := newError(http.StatusInternalServerError, err.Error())
		render.JSON(w, r, err)
		return
	}

	render.JSON(w, r, convertUser(user))
}

func convertUser(user *entity.User) *generated.User {
	if user == nil {
		return nil
	}

	return &generated.User{
		Id:   &user.ID,
		Name: &user.Name,
	}
}
