package middleware

import (
	user "car-api/internal/core/domain"
	"car-api/internal/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"time"
)

func JwtMiddleware(tokenKey string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(tokenKey),
	})
}

func SignToken(tokenKey string, user *user.UserResponse) string {

	// Create the Claims
	claims := jwt.MapClaims{
		"admin":    true,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
		"username": user.Username,
		"id":       user.ID,
		"role":     user.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ""
	}

	return t
	//token := jwt.New(jwt.SigningMethodHS256)
	//claims := token.Claims.(jwt.MapClaims)
	//claims["admin"] = true
	//claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	//claims["sub"] = id
	//
	//t, err := token.SignedString([]byte(tokenKey))
	//
	//if err != nil {
	//	return ""
	//}
	//
	//return t
}

func ExtractUserIDFromJWT(bearer, tokenKey string) string {

	token := bearer[7:]
	logs.Info(token)
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenKey), nil
	})

	if err != nil {
		return ""
	}

	if t.Valid {
		claims := t.Claims.(jwt.MapClaims)
		return claims["sub"].(string)
	}

	return ""
}
