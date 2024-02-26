package controller

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/koki-algebra/go_server_sample/internal/entity"
	"github.com/koki-algebra/go_server_sample/internal/infra/http/oapi"
)

func (ctrl *controllerImpl) GetUserByID(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()
	user, err := ctrl.user.GetByID(ctx, id)
	if err != nil {
		// todo: error handling
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, ParseError(err))
	}

	render.JSON(w, r, convertUser(user))
}

func convertUser(user *entity.User) *oapi.User {
	if user == nil {
		return nil
	}

	return &oapi.User{
		Id:   &user.ID,
		Name: &user.Name,
	}
}
