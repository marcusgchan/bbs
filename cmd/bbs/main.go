package main

import (
	"github.com/labstack/echo/v4"
	"github.com/marcusgchan/bbs/database"
	"github.com/marcusgchan/bbs/internal/testEvents"
)

func main() {
	app := echo.New()

	testEventsGroup := app.Group("/test-events")
	testEvents.UseTestEventRoutes(testEventsGroup)

	api := app.Group("/api")
	bootStrapApiRoutes(api)

	// app.Start(":3000")
	database.Seed()
}

func bootStrapApiRoutes(g *echo.Group) {
	testEventsGroup := g.Group("/test-events")
	testEvents.UseInjestTestEventRoutes(testEventsGroup)
}
