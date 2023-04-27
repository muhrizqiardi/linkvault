package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"
	"server/handlers"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
)

func JwtAuth(ctx context.Context, l *log.Logger, pg *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			token := strings.Split(authHeader, " ")
			if len(token) == 1 {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			parsedToken, err := jwt.ParseWithClaims(token[1], &handlers.Claims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claims, ok := parsedToken.Claims.(*handlers.Claims)
			if !ok && !parsedToken.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			nextCtx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(nextCtx))
		})
	}
}