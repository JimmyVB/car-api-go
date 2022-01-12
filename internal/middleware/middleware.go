package middleware

import (
	user "car-api/internal/core/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

func JwtMiddleware(tokenKey string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		headers := c.GetReqHeaders()
		strToken := headers["Authorization"]
		if strToken == "" {
			return c.SendStatus(fiber.StatusForbidden)
		}
		splitToken := strings.Split(strToken, "Bearer ")
		reqToken := splitToken[1]
		t, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(tokenKey), nil
		})
		if err != nil {
			return fiber.NewError(fiber.StatusForbidden, "token invalido")
		}
		if t.Valid {
			claims := t.Claims.(jwt.MapClaims)
			role := claims["role"].(string)
			if role != "ROLE_ADMIN" {
				return fiber.NewError(fiber.StatusUnauthorized, "no tiene privilegios para ejecutar esta accion")
			} else {
				return c.Next()
			}
		} else if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return fiber.NewError(fiber.StatusForbidden, "token invalido")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return fiber.NewError(fiber.StatusUnauthorized, "token expirado")
			} else {
				return fiber.NewError(fiber.StatusForbidden, "token invalido")
			}
		}
		return c.Next()
	}
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

func ExtractUserIDFromJWT(token, tokenKey string) string {
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
