package testEvents

import (
	"github.com/labstack/echo/v4"
)

func UseTestEventRoutes(g *echo.Group) {
	g.GET("", handleShowTestEvents)
}

func UseInjestTestEventRoutes(g *echo.Group) {
	g.POST("", handleInjestTestEvent)
}
