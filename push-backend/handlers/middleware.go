package handlers

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userContextKey = contextKey("user")

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.ParseWithClaims(
			tokenString,
			jwt.MapClaims{},
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
					return nil, jwt.ErrTokenUnverifiable
				}
				return jwtSecret, nil
			},
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		)
		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, token.Claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func UserContextKey() interface{} {
	return userContextKey
}

func GetUserIDFromContext(r *http.Request) (int64, error) {
	claimsVal := r.Context().Value(UserContextKey())
	claims, ok := claimsVal.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("could not extract user claims")
	}
	idVal, ok := claims["id"]
	if !ok {
		return 0, errors.New("user ID missing in token")
	}
    userID, ok := idVal.(float64)
    if !ok {
        return 0, errors.New("user ID in token has wrong type")
    }
	return int64(userID), nil
}
