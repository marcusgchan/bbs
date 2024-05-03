package auth

import (
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/marcusgchan/bbs/internal"
	auth "github.com/marcusgchan/bbs/internal/auth/views"
)

func Authenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		accessToken, err := c.Cookie("access-token")
		if err != nil {
			return internal.Render(auth.LoginPage(), c)
		}

		token, err := jwt.Parse(accessToken.Value, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return internal.Render(auth.LoginPage(), c)
		}

		return next(c)
	}
}

func ApiAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h := c.Request().Header.Get("Authorization")
		ah := strings.Split(h, " ")
		if len(ah) != 2 {
			return fmt.Errorf("invalid authorization header format")
		}
		key := ah[1]
		if key != os.Getenv("API_KEY") {
			return fmt.Errorf("invalid api key")
		}
		return next(c)
	}
}
