package handler

import "github.com/koki-algebra/go_server_sample/internal/infra/http/generated"

func newError(code int, message string) generated.Error {
	return generated.Error{
		Code:    int32(code),
		Message: message,
	}
}
