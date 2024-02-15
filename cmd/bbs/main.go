package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/marcusgchan/bbs/database"
	slqc "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal/auth"
	"github.com/marcusgchan/bbs/internal/player"
	"github.com/marcusgchan/bbs/internal/testevt"
)

func main() {
	app := echo.New()
	db := database.Connect()
	q := slqc.New(db)

	app.Use(middleware.Logger())
	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "build",
		Browse: false,
	}))

	app.GET("/login", auth.AuthHandler{Q: q, DB: db}.Login)
	app.POST("/login", auth.AuthHandler{Q: q, DB: db}.HandleLogin)

	testEventsGroup := app.Group("/test-events")
	testEventsGroup.Use(auth.Authenticated)
	testEventsHandler := testevt.TestEventHandler{Q: q, DB: db}
	testEventsGroup.GET("", testEventsHandler.ShowTestEvent)

	playersGroup := app.Group("/players")
	playersGroup.Use(auth.Authenticated)
	playersHandler := player.PlayerHandler{Q: q, DB: db}
	playersGroup.GET("", playersHandler.ShowPlayerList)

	// Not found
	app.Any("/*", func(c echo.Context) error {
		return c.String(404, "Page not found.")
	})

	api := app.Group("/api")
	// Remember to auth!!!!!!!!!!!!!!!!!!!!!!!!!
	testEvtApiGroup := api.Group("/test-events")
	testEvtApiGroup.POST("", testEventsHandler.CreateTestEvent)

	app.Start(":3000")
}
