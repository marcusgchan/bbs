package stats

import (
	"database/sql"
	"strconv"

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
	limit, err := strconv.Atoi(c.QueryParam("numberOfVersions"))
	if err != nil {
		limit = 3
	}
	version := c.QueryParam("version")
	var singleStatsResP *database.GetStatsByVersionRow
	if version == "" {
		singleStatsRes, err := h.Q.GetStatsByVersion(c.Request().Context(), database.GetStatsByVersionParams{
			Version:   version,
			Version_2: version,
			Version_3: version,
			Version_4: version,
			Version_5: version,
		})
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if err != sql.ErrNoRows {
			singleStatsResP = &singleStatsRes
		}
	}
	muliStatsRes, err := h.Q.GetMostRecentStats(c.Request().Context(), database.GetMostRecentStatsParams{
		Limit:   int64(limit),
		Limit_2: int64(limit),
		Limit_3: int64(limit),
		Limit_4: int64(limit),
		Limit_5: int64(limit),
	})
	if err != nil {
		return err
	}

	if internal.FromHTMX(c) {
		return internal.Render(stats.StatsContent(TransformToStatsProps(singleStatsResP, &muliStatsRes)), c)
	}
	return internal.Render(stats.StatsPage(TransformToStatsProps(singleStatsResP, &muliStatsRes)), c)
}
