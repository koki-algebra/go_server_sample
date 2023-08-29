package handler

import (
	"context"

	"github.com/koki-algebra/grpc_sample/internal/infra/grpc/generated"
)

type UserHandler struct {
	generated.UserServiceServer
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetByID(ctx context.Context, req *generated.GetByIDRequest) (*generated.GetByIDResponse, error) {
	return &generated.GetByIDResponse{
		Id:   "xyz",
		Name: "koki",
	}, nil
}
