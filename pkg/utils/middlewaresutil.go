package utils

import "net/http"

// middleware is a function that hhtp.Handler with additional Functanality
type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
