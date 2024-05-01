package stats

import (
	"database/sql"
	"strconv"

	"github.com/labstack/echo/v4"
	database "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal"
	"github.com/marcusgchan/bbs/internal/stats/views"
	"github.com/marcusgchan/bbs/internal/sview"
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
	var singleStatsRes database.GetStatsByVersionRow
	if version != "" {
		singleStatsRes, err = h.Q.GetStatsByVersion(c.Request().Context(), database.GetStatsByVersionParams{
			Version:   version,
			Version_2: version,
			Version_3: version,
			Version_4: version,
			Version_5: version,
		})
		if err != nil && err != sql.ErrNoRows {
			return err
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

	versions, err := h.Q.GetVersions(c.Request().Context())
	if err != nil {
		return err
	}

	statsProps := stats.StatsPageProps{
		Single:   TransformToSingleField(&singleStatsRes),
		Multi:    TransformToMultiField(&muliStatsRes),
		Versions: TransformToVersionsField(&versions),
	}

	if internal.FromHTMX(c) {
		return internal.Render(stats.StatsContent(&statsProps), c)
	}
	return internal.Render(stats.StatsPage(&statsProps), c)
}

func (h StatsHandler) LatestVersions(c echo.Context) error {
	if !internal.FromHTMX(c) {
		return internal.Render(sview.NotFoundPage(), c)
	}
	numOfVersions, err := strconv.Atoi(c.QueryParam("numverOfVersions"))
	if err != nil || numOfVersions < 3 || numOfVersions > 10 {
		numOfVersions = 3
	}
	data, err := h.Q.GetMostRecentStats(c.Request().Context(), database.GetMostRecentStatsParams{
		Limit:   int64(numOfVersions),
		Limit_2: int64(numOfVersions),
		Limit_3: int64(numOfVersions),
		Limit_4: int64(numOfVersions),
		Limit_5: int64(numOfVersions),
		Limit_6: int64(numOfVersions),
	})
	if err != nil {
		return err
	}
	return internal.Render(stats.RecentStatsList(TransformToMultiField(&data)), c)
}

func (h StatsHandler) FilteredStats(c echo.Context) error {
	if !internal.FromHTMX(c) {
		return internal.Render(sview.NotFoundPage(), c)
	}
	version := c.QueryParam("version")
	data, err := h.Q.GetStatsByVersion(c.Request().Context(), database.GetStatsByVersionParams{
		Version: version,
	})
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return internal.Render(stats.FilteredStats(TransformToStatsField(&data)), c)
}
