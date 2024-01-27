package testEvents

import (
	"github.com/labstack/echo/v4"
	"github.com/marcusgchan/bbs/internal/testEvents/views"
)

func UseTestEventRoutes(g *echo.Group) {
	g.GET("", handleShowTestEvents)
}

func handleShowTestEvents(c echo.Context) error {
	return views.Page().Render(c.Request().Context(), c.Response())
}
