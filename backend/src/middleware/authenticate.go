package middleware

import (
	"context"
	"net/http"

	"github.com/muskiteer/chat-app/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type contextKey string

const UserIDKey contextKey = "userId"

// Authenticate middleware validates JWT and attaches userId to context
func Authenticate(userCollection *mongo.Collection) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("jwt")
			if err != nil {
				utils.JSONError(w, http.StatusUnauthorized, "Unauthorized: No token provided")
				return
			}

			claims, err := utils.VerifyJWT(cookie.Value, userCollection)
			if err != nil {
				utils.JSONError(w, http.StatusUnauthorized, "err.Error()")
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
