package testEvents

import (
	"github.com/labstack/echo/v4"
	"github.com/marcusgchan/bbs/internal"
)

func UseTestEventRoutes(ctx *internal.Context) {
	g.GET("", handleShowTestEvents)
}

func UseInjestTestEventRoutes(g *echo.Group) {
	g.POST("", handleInjestTestEvent)
}
