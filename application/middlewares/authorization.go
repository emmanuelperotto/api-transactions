package middlewares

import (
	"log"
	"net/http"
	"strings"
	"transactions/application/authentication"
)

var (
	CurrentAccountID int
)

func AuthorizeRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		accessToken := strings.ReplaceAll(authHeader, "Bearer ", "")

		claims, err := authentication.JsonWebToken.Decode(accessToken)

		if err != nil {
			log.Println("[JWT ERROR]", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		CurrentAccountID = int(claims["id"].(float64))

		next.ServeHTTP(w, r)
	})
}
