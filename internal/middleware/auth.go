package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/ilyushkaaa/banner-service/internal/banner/delivery"
	"github.com/ilyushkaaa/banner-service/internal/pkg/response"
	"github.com/ilyushkaaa/banner-service/internal/user/storage"
)

func (mw *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Token")
		user, err := mw.userService.GetUserByToken(r.Context(), token)
		if err != nil {
			if errors.Is(err, storage.ErrUserNotFound) {
				mw.logger.Errorf("no users with token %s", token)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			mw.logger.Errorf("internal server error in authorization: %v", err)
			response.WriteResponse(w, response.Error{Err: response.ErrInternal.Error()}, http.StatusInternalServerError, mw.logger)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, delivery.UserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
