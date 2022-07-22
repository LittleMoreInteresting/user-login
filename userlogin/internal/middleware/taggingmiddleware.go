package middleware

import (
	"context"
	"net/http"
)

type TaggingMiddleware struct {
}

func NewTaggingMiddleware() *TaggingMiddleware {
	return &TaggingMiddleware{}
}

func (m *TaggingMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "tag", "tagV111")
		next(w, r.WithContext(ctx))
	}
}
