package auth

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	database "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal"
	auth "github.com/marcusgchan/bbs/internal/auth/views"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Q  *database.Queries
	DB *sql.DB
}

func (h AuthHandler) Login(c echo.Context) error {
	_, err := c.Cookie("access-token")
	if err != nil {
		return internal.Render(auth.LoginPage(), c)
	}
	if strings.Contains(c.Request().URL.String(), "/login") {
		c.Redirect(302, "/test-events")
	}
	return c.Redirect(302, c.Request().Header.Get("HX-CURRENT-URL"))
}

func (h AuthHandler) HandleLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// User does not exist
	ctx := context.Background()
	user, err := h.Q.GetUser(ctx, username)
	if err != nil {
		return internal.Render(auth.Error(), c)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return internal.Render(auth.Error(), c)
	}

	expDate := time.Now().Add(time.Hour * 24 * 7)
	claims := jwt.MapClaims{
		"iss": "bbs",
		"aud": "admin",
		"sub": user.Username,
		"iat": time.Now().Unix(),
		"exp": expDate.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return internal.Render(auth.Error(), c)
	}

	c.SetCookie(&http.Cookie{
		Name:     "access-token",
		Value:    tokenStr,
		Expires:  expDate,
		SameSite: http.SameSiteStrictMode,
	})

	headers := c.Response().Header()
	headers.Set("HX-Refresh", "true")
	return c.NoContent(200)
}
