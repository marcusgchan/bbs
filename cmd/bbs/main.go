package main

import (
	"github.com/labstack/echo/v4"
	"github.com/marcusgchan/bbs/internal/testEvents"
)

func main() {
	app := echo.New()

	testEventsGroup := app.Group("/test-events")
	testEvents.UseTestEventRoutes(testEventsGroup)

	app.Start(":3000")
}
