package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/marcusgchan/bbs/database"
	slqc "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal"
	"github.com/marcusgchan/bbs/internal/auth"
	"github.com/marcusgchan/bbs/internal/player"
	"github.com/marcusgchan/bbs/internal/stats"
	"github.com/marcusgchan/bbs/internal/sview"
	"github.com/marcusgchan/bbs/internal/template"
	"github.com/marcusgchan/bbs/internal/testevt"
	"github.com/marcusgchan/bbs/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}
	app := echo.New()
	db := database.Connect()
	q := slqc.New(db)

	// app.Use(middleware.Logger())
	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(web.StaticFS),
	}))

	app.GET("/login", auth.AuthHandler{Q: q, DB: db}.Login)
	app.POST("/login", auth.AuthHandler{Q: q, DB: db}.HandleLogin)

	testEventsGroup := app.Group("/test-events")
	testEventsGroup.Use(auth.Authenticated)
	testEventsHandler := testevt.TestEventHandler{Q: q, DB: db}
	testEventsGroup.GET("", testEventsHandler.TestEvtPage)
	testEventsGroup.GET("/:testEventId", testEventsHandler.TestEvtResPage)

	playersGroup := app.Group("/players")
	playersGroup.Use(auth.Authenticated)
	playersHandler := player.PlayerHandler{Q: q, DB: db}
	playersGroup.GET("", playersHandler.PlayerListPage)
	playersGroup.GET("/:playerId", playersHandler.PlayerInfoPage)

	statsGroup := app.Group("/stats")
	statsGroup.Use(auth.Authenticated)
	statsHandler := stats.StatsHandler{Q: q, DB: db}
	statsGroup.GET("", statsHandler.StatsPage)
	statsGroup.GET("/latest-versions", statsHandler.LatestVersions)
	statsGroup.GET("/filtered", statsHandler.FilteredStats)

	app.GET("/*", func(c echo.Context) error {
		return internal.Render(sview.NotFoundPage(), c)
	}, auth.Authenticated)

	api := app.Group("/api")
	// Remember to auth!!!!!!!!!!!!!!!!!!!!!!!!!
	testEvtApiGroup := api.Group("/test-events")
	testEvtApiGroup.POST("", testEventsHandler.CreateTestEvent)
	testEvtApiGroup.POST("/results", testEventsHandler.CreatePlayerTestResult)

	playerApiGroup := api.Group("/players")
	playerApiGroup.POST("", playersHandler.CreatePlayer)

	templateApiGroup := api.Group("/templates")
	templateHandler := template.TemplateHandler{Q: q, DB: db}
	templateApiGroup.POST("", templateHandler.Create)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Start(":" + port)
}
