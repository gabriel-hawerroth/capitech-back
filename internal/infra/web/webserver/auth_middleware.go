package webserver

import (
	"net/http"
	"strings"
)

var (
	// secretKey   = []byte("secret-key")
	allowedUrls = []string{}
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		for _, url := range allowedUrls {
			if strings.ContainsAny(url, path) {
				next.ServeHTTP(w, r)
				return
			}
		}

		// authorizationHeader := r.Header.Get("Authorization")
		// if authorizationHeader == "" {
		// 	http.Error(w, "Token de autorizacao nao fornecido", http.StatusUnauthorized)
		// 	return
		// }

		// tokenString := authorizationHeader[len("Bearer "):]

		// token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 	return secretKey, nil
		// })
		// if err != nil || !token.Valid {
		// 	http.Error(w, "Token de autorização inválido", http.StatusUnauthorized)
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}
