package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/marcusgchan/bbs/database"
	slqc "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal"
	"github.com/marcusgchan/bbs/internal/auth"
	"github.com/marcusgchan/bbs/internal/player"
	"github.com/marcusgchan/bbs/internal/sview"
	"github.com/marcusgchan/bbs/internal/testevt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}
	app := echo.New()
	db := database.Connect()
	q := slqc.New(db)

	app.Use(middleware.Logger())
	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "build",
		Browse: true,
	}))

	app.GET("/login", auth.AuthHandler{Q: q, DB: db}.Login)
	app.POST("/login", auth.AuthHandler{Q: q, DB: db}.HandleLogin)

	testEventsGroup := app.Group("/test-events")
	testEventsGroup.Use(auth.Authenticated)
	testEventsHandler := testevt.TestEventHandler{Q: q, DB: db}
	testEventsGroup.GET("", testEventsHandler.GetTestEvtPage)
	testEventsGroup.GET("/content", testEventsHandler.GetTestEvtContent)

	playersGroup := app.Group("/players")
	playersGroup.Use(auth.Authenticated)
	playersHandler := player.PlayerHandler{Q: q, DB: db}
	playersGroup.GET("", playersHandler.ShowPlayerList)

	// Not found
	app.Any("/*", func(c echo.Context) error {
		return internal.Render(sview.Base(), c)
	}, auth.Authenticated)

	api := app.Group("/api")
	// Remember to auth!!!!!!!!!!!!!!!!!!!!!!!!!
	testEvtApiGroup := api.Group("/test-events")
	testEvtApiGroup.POST("", testEventsHandler.CreateTestEvent)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Start(":" + port)
}
