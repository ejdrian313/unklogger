package infrastrucutre

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var error401 = map[string]string{
	"error": "Unauthorized",
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headerValue := c.Request().Header.Get("Authorization")
		splits := strings.Split(headerValue, " ")
		if len(splits) != 2 || splits[0] != "Bearer" {
			return c.JSON(401, error401)
		}
		rawToken := splits[1]
		token, err := validateJWT(rawToken)
		c.Set("username", token.Claims.(jwt.MapClaims)["username"].(string))
		c.Set("ID", uint(token.Claims.(jwt.MapClaims)["ID"].(float64)))
		if err != nil {
			return c.JSON(401, error401)
		}

		return next(c)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}
