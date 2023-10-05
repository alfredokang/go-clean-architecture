package middlewares

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
	"strings"
	"time"
)

type AuthMiddleware struct {
	Token *jwtauth.JWTAuth
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{
		Token: jwtauth.New("HS512", []byte("uhu"), nil),
	}
}

var jwtKey = []byte("chave_secreta_do_jwt")

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims, err := VerifyToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Adicione os dados do usuário aos contextos da requisição, se necessário
		ctx := context.WithValue(r.Context(), "claims", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
