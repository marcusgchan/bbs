package testEvent

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/marcusgchan/bbs/internal"
	"github.com/marcusgchan/bbs/internal/testEvent/views"
)

type TestEventHandler struct {
	DB *sql.DB
}

func (h TestEventHandler) ShowTestEvent(c echo.Context) error {
	return internal.Render(testEvent.Page(), c)
}
