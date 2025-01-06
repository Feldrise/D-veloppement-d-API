package authentication

import (
	"context"
	"net/http"

	"feldrise.com/animal-api/config"
	"feldrise.com/animal-api/database/dbmodel"
)

var UserCtxKey = contextKey{"user"}

type contextKey struct {
	name string
}

func Middelware(c *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			// We get the user from the token
			tokenString := header
			userID, _, err := ParseToken(c.Constants.JWTSecret, tokenString)

			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			user, err := c.UserRepository.FindByID(userID, true)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// We add the user to the context
			ctx := context.WithValue(r.Context(), UserCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := ForContext(r.Context())

		if user == nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func ForContext(ctx context.Context) *dbmodel.User {
	raw, _ := ctx.Value(UserCtxKey).(*dbmodel.User)
	return raw
}
