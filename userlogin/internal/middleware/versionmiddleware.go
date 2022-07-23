package middleware

import (
	"context"
	"net/http"
)

type VersionMiddleware struct {
}

func NewVersionMiddleware() *VersionMiddleware {
	return &VersionMiddleware{}
}

func (m *VersionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "version", "v1.1.0")
		next(w, r.WithContext(ctx))
	}
}
