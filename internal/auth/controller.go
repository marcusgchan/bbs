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
	return c.Redirect(302, "/")
}

func (h AuthHandler) HandleLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	type User struct {
		ID       int
		Username string
		Password string
	}

	user := User{}

	err := h.DB.QueryRow("SELECT * FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	defer h.DB.Close()
	if err != nil {
		return c.String(401, "Unauthorized")
	}

	if user.Password != password {
		return c.String(401, "Unauthorized")
	}

	return c.Redirect(302, "/")
}
