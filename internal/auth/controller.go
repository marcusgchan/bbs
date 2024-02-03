package auth

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"unicode/utf8"

	sq "github.com/Masterminds/squirrel"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/marcusgchan/bbs/internal"
	"github.com/marcusgchan/bbs/internal/auth/views"
	"github.com/marcusgchan/bbs/internal/testEvent/views"
	"golang.org/x/crypto/bcrypt"
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
	return c.Redirect(302, "/test-events")
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

	err := sq.
		Select("username, password").
		RemoveOffset().Limit(1).
		From("users").
		Where(sq.Eq{"username": username}).
		RunWith(h.DB).
		QueryRow().
		Scan(&user.Username, &user.Password)
		// User does not exist
	if err != nil {
		log.Print(err.Error())
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
		log.Print(err.Error())
		return c.String(500, "Internal Server Error")
	}

	c.SetCookie(&http.Cookie{
		Name:    "access-token",
		Value:   tokenStr,
		Expires: expDate,
	})

	headers := c.Response().Header()
	headers.Set("HX-Retarget", "main")

	callbackUrl := c.QueryParam("callback")
	if utf8.RuneCountInString(callbackUrl) == 0 {
		headers.Set("HX-Push-Url", "/test-events")
		return internal.Render(testEvent.Page(), c)
	}

	return c.Redirect(302, callbackUrl)
}
