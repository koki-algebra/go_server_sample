package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler

func With(h http.Handler, ms ...Middleware) http.Handler {
	for _, m := range ms {
		h = m(h)
	}

	return h
}
