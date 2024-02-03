package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/marcusgchan/bbs/database"
	"github.com/marcusgchan/bbs/internal/auth"
	"github.com/marcusgchan/bbs/internal/testEvent"
)

func main() {
	app := echo.New()
	db := database.Connect()

	// app.Use(middleware.Logger())
	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "build",
		Browse: false,
	}))

	app.GET("/login", auth.AuthHandler{DB: db}.Login)
	app.POST("/login", auth.AuthHandler{DB: db}.HandleLogin)

	testEventsGroup := app.Group("/test-events")
	testEventsGroup.Use(auth.Authenticated)
	testEventsHandler := testEvent.TestEventHandler{DB: db}
	testEventsGroup.GET("", testEventsHandler.ShowTestEvent)

	// api := app.Group("/api")
	// bootStrapApiRoutes(api)

	app.Start(":3000")
}

// func bootStrapApiRoutes(g *echo.Group) {
// 	testEventsGroup := g.Group("/test-events")
// 	testEvents.UseInjestTestEventRoutes(testEventsGroup)
// }
