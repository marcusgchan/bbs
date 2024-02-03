package auth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Authenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken, err := c.Cookie("access-token")
		reqUrl := c.Request().Host
		callBackUrl := fmt.Sprintf("/login?callback=%v", reqUrl)
		if err != nil {
			return c.Redirect(302, callBackUrl)
		}

		token, err := jwt.Parse(accessToken.Value, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return c.Redirect(401, callBackUrl)
		}

		return next(c)
	}
}
