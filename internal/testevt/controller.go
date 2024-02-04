package testevt

import (
	"github.com/labstack/echo/v4"
	"github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal"
	"github.com/marcusgchan/bbs/internal/testevt/views"
)

type TestEventHandler struct {
	DB *database.Queries
}

func (h TestEventHandler) ShowTestEvent(c echo.Context) error {
	return internal.Render(testevt.Page(), c)
}
