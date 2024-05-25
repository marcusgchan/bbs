package stats

import (
	"database/sql"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
	database "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal"
	stats "github.com/marcusgchan/bbs/internal/stats/views"
	"github.com/marcusgchan/bbs/internal/sview"
)

type StatsHandler struct {
	Q  *database.Queries
	DB *sql.DB
}

func (h StatsHandler) StatsPage(c echo.Context) error {
	reqType := c.Request().Header.Get("HX-Trigger-Name")
	var err error
	switch reqType {
	case "numberOfVersionsForTestEvent":
		err = partialStatsPageReq(h, c)
	// case "numberOfVersionsForCatastrophe":
	// 	err = catastrophe(h, c)
	// case "version":
	// 	err = testEvent(h, c)
	default:
		err = normalStatsPageReq(h, c)
	}
	return err
}

func partialStatsPageReq(h StatsHandler, c echo.Context) error {
	if !internal.FromHTMX(c) {
		return internal.Render(sview.NotFoundPage(), c)
	}
	limit := parseTestEventQueryParams(c.QueryParam("numberOfVersionsForTestEvent"))

	err := SetPushUrlInHeader(c, strconv.Itoa(limit))
	if err != nil {
		return err
	}

	normalTestEvtStats, err := h.Q.GetTestEventsStats(c.Request().Context(), database.GetTestEventsStatsParams{
		Limit:        int64(limit),
		Limit_2:      int64(limit),
		Limit_3:      int64(limit),
		Limit_4:      int64(limit),
		Limit_5:      int64(limit),
		Limit_6:      int64(limit),
		Difficulty:   "normal",
		Difficulty_2: "normal",
		Difficulty_3: "normal",
		Difficulty_4: "normal",
		Difficulty_5: "normal",
		Difficulty_6: "normal",
	})
	if err != nil {
		return err
	}
	hardTestEvtStats, err := h.Q.GetTestEventsStats(c.Request().Context(), database.GetTestEventsStatsParams{
		Limit:        int64(limit),
		Limit_2:      int64(limit),
		Limit_3:      int64(limit),
		Limit_4:      int64(limit),
		Limit_5:      int64(limit),
		Limit_6:      int64(limit),
		Difficulty:   "hard",
		Difficulty_2: "hard",
		Difficulty_3: "hard",
		Difficulty_4: "hard",
		Difficulty_5: "hard",
		Difficulty_6: "hard",
	})
	if err != nil {
		return err
	}

	catData, err := h.Q.GetCatastropheKills(c.Request().Context(), int64(limit))
	if err != nil {
		return err
	}

	statsByVersion := mergeGeneralStats(&normalTestEvtStats, &hardTestEvtStats)
	addCatastropheDeathsToStats(statsByVersion, &catData)

	return internal.Render(stats.StatsByVersionList(statsByVersion), c)
}

func SetPushUrlInHeader(c echo.Context, val string) error {
	url, err := url.Parse(c.Request().Header.Get("HX-Current-Url"))
	if err != nil {
		return err
	}
	qp := url.Query()
	keyToReplace := c.Request().Header.Get("HX-Trigger-Name")
	qp.Del(keyToReplace)
	qp.Add(keyToReplace, val)
	c.Response().Header().Set("HX-Push-Url", url.Path+"?"+qp.Encode())
	return nil
}

// func testEvent(h StatsHandler, c echo.Context) error {
// 	if !internal.FromHTMX(c) {
// 		return internal.Render(sview.NotFoundPage(), c)
// 	}
// 	version := c.QueryParam("version")
// 	data, err := h.Q.GetTestEventStatsByVersion(c.Request().Context(), database.GetTestEventStatsByVersionParams{
// 		Version:   version,
// 		Version_2: version,
// 		Version_3: version,
// 		Version_4: version,
// 		Version_5: version,
// 		Version_6: version,
// 	})
// 	if err != nil && err != sql.ErrNoRows {
// 		return err
// 	}
//
// 	err = SetPushUrlInHeader(c, version)
// 	if err != nil {
// 		return err
// 	}
//
// 	return internal.Render(stats.FilteredStats(TransformToStatsField(&data)), c)
// }

// func catastrophe(h StatsHandler, c echo.Context) error {
// 	if !internal.FromHTMX(c) {
// 		return internal.Render(sview.NotFoundPage(), c)
// 	}
//
// 	count := parseCatastropheQueryParams(c.QueryParam("numberOfVersionsForCatastrophe"))
//
// 	data, err := h.Q.GetCatastropheKills(c.Request().Context(), int64(count))
// 	if err != nil {
// 		return err
// 	}
//
// 	err = SetPushUrlInHeader(c, strconv.Itoa(count))
// 	if err != nil {
// 		return err
// 	}
//
// 	return internal.Render(stats.CatastropheStatsList(TransformToCatastropheField(&data)), c)
// }

func normalStatsPageReq(h StatsHandler, c echo.Context) error {
	limit := parseTestEventQueryParams(c.QueryParam("numberOfVersionsForTestEvent"))
	version := c.QueryParam("version")
	singleStatsRes := database.GetTestEventStatsByVersionRow{}
	var err error
	if version != "" {
		singleStatsRes, err = h.Q.GetTestEventStatsByVersion(c.Request().Context(), database.GetTestEventStatsByVersionParams{
			Version:   version,
			Version_2: version,
			Version_3: version,
			Version_4: version,
			Version_5: version,
			Version_6: version,
		})
		if err != nil && err != sql.ErrNoRows {
			return err
		}
	}
	versions, err := h.Q.GetVersions(c.Request().Context())
	if err != nil {
		return err
	}
	normalTestEvtStats, err := h.Q.GetTestEventsStats(c.Request().Context(), database.GetTestEventsStatsParams{
		Limit:        int64(limit),
		Limit_2:      int64(limit),
		Limit_3:      int64(limit),
		Limit_4:      int64(limit),
		Limit_5:      int64(limit),
		Limit_6:      int64(limit),
		Difficulty:   "normal",
		Difficulty_2: "normal",
		Difficulty_3: "normal",
		Difficulty_4: "normal",
		Difficulty_5: "normal",
		Difficulty_6: "normal",
	})
	if err != nil {
		return err
	}
	hardTestEvtStats, err := h.Q.GetTestEventsStats(c.Request().Context(), database.GetTestEventsStatsParams{
		Limit:        int64(limit),
		Limit_2:      int64(limit),
		Limit_3:      int64(limit),
		Limit_4:      int64(limit),
		Limit_5:      int64(limit),
		Limit_6:      int64(limit),
		Difficulty:   "hard",
		Difficulty_2: "hard",
		Difficulty_3: "hard",
		Difficulty_4: "hard",
		Difficulty_5: "hard",
		Difficulty_6: "hard",
	})
	if err != nil {
		return err
	}

	catData, err := h.Q.GetCatastropheKills(c.Request().Context(), int64(limit))
	if err != nil {
		return err
	}

	defaults := &stats.InputDefaults{
		RecentTestEvents: strconv.Itoa(limit),
		TestEvent:        version,
	}

	if singleStatsRes != (database.GetTestEventStatsByVersionRow{}) {
		defaults.TestEvent = version
	}

	statsByVersion := mergeGeneralStats(&normalTestEvtStats, &hardTestEvtStats)
	addCatastropheDeathsToStats(statsByVersion, &catData)

	statsProps := stats.StatsPageProps{
		Versions:       TransformToVersionsField(&versions),
		StatsByVersion: statsByVersion,
		InputDefaults:  defaults,
	}

	if internal.FromHTMX(c) {
		return internal.Render(stats.StatsContent(&statsProps), c)
	}
	return internal.Render(stats.StatsPage(&statsProps), c)
}
