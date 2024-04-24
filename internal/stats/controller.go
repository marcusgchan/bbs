package stats

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	database "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal"
	"github.com/marcusgchan/bbs/internal/stats/views"
)

type StatsHandler struct {
	Q  *database.Queries
	DB *sql.DB
}

func (h StatsHandler) StatsPage(c echo.Context) error {
	if internal.FromHTMX(c) {
		return internal.Render(stats.StatsContent(), c)
	}
	return internal.Render(stats.StatsPage(), c)
}
