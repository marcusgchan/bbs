package auth

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/marcusgchan/bbs/internal"
	"github.com/marcusgchan/bbs/internal/auth/views"
)

type AuthHandler struct {
	DB *sql.DB
}

func (h AuthHandler) Login(c echo.Context) error {
	token, err := c.Cookie("access-token")
	if err != nil {
		return internal.Render(auth.LoginPage(), c)
	}
	fmt.Printf("token: %s\n", token)
	return nil
}
