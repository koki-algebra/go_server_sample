package controller

import "github.com/koki-algebra/go_server_sample/internal/infra/http/oapi"

func ParseError(err error) oapi.Error {
	return oapi.Error{
		Message: err.Error(),
	}
}
