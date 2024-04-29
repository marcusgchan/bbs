package stats

import (
	"fmt"
	"strconv"

	database "github.com/marcusgchan/bbs/database/gen"
	stats "github.com/marcusgchan/bbs/internal/stats/views"
)

func TransformToStatsProps(singleStats *database.GetStatsByVersionRow, multiStats *[]database.GetMostRecentStatsRow) *stats.StatsPageProps {
	multiStatsProps := make([]stats.Stats, len(*multiStats))
	for i, s := range *multiStats {
		multiStatsProps[i] = stats.Stats{
			Version:         s.Version,
			StartDate:       s.Startdate,
			EndDate:         s.Enddate,
			AvgWaveSurvived: fmt.Sprintf("%f.2", s.Avgwave),
			HighestWave:     strconv.Itoa(int(s.Maxwave)),
			Count:           strconv.Itoa(int(s.Numoftestevents)),
		}
	}
	if singleStats == nil {
		return &stats.StatsPageProps{
			Multi: &multiStatsProps,
		}
	}
	return &stats.StatsPageProps{
		Multi: &multiStatsProps,
		Single: &stats.Stats{
			Version:         singleStats.Version,
			StartDate:       singleStats.Startdate.String(),
			EndDate:         singleStats.Enddate.String(),
			AvgWaveSurvived: fmt.Sprintf("%.2f", singleStats.Avgwave.Float64),
			HighestWave:     strconv.Itoa(int(singleStats.Maxwave)),
			Count:           strconv.Itoa(int(singleStats.Numoftestevents)),
		},
	}
}
