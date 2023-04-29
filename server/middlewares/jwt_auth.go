package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"
	"server/handlers"
	"server/utils"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
)

func JwtAuth(ctx context.Context, l *log.Logger, pg *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.BaseResponseWriter[any](w, http.StatusUnauthorized, false, "Unauthorized", nil)
				l.Println("Unauthorized request: empty Authorization header")
				return
			}

			token := strings.Split(authHeader, " ")
			if len(token) == 1 {
				utils.BaseResponseWriter[any](w, http.StatusUnauthorized, false, "Unauthorized", nil)
				l.Println("Unauthorized request: invalid Authorization header")
			}

			parsedToken, err := jwt.ParseWithClaims(token[1], &handlers.Claims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})
			if err != nil {
				utils.BaseResponseWriter[any](w, http.StatusUnauthorized, false, "Unauthorized", nil)
				l.Println(err.Error())
			}

			claims, ok := parsedToken.Claims.(*handlers.Claims)
			if !ok && !parsedToken.Valid {
				utils.BaseResponseWriter[any](w, http.StatusUnauthorized, false, "Unauthorized", nil)
				l.Println("Failed to assert type")
			}

			nextCtx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(nextCtx))
		})
	}
}