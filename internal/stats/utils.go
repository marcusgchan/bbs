package stats

import (
	"fmt"
	"strconv"

	database "github.com/marcusgchan/bbs/database/gen"
	stats "github.com/marcusgchan/bbs/internal/stats/views"
)

func TransformSingleAndMultiToStatsProps(singleStats *database.GetStatsByVersionRow, multiStats *[]database.GetMostRecentStatsRow) *stats.StatsPageProps {
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
			StartDate:       singleStats.Startdate,
			EndDate:         singleStats.Enddate,
			AvgWaveSurvived: fmt.Sprintf("%.2f", singleStats.Avgwave),
			HighestWave:     strconv.Itoa(int(singleStats.Maxwave)),
			Count:           strconv.Itoa(int(singleStats.Numoftestevents)),
		},
	}
}

func TransformMultiToStats(data *[]database.GetMostRecentStatsRow) *[]stats.Stats {
	s := make([]stats.Stats, len(*data))
	for i, v := range *data {
		s[i] = stats.Stats{
			Version:         v.Version,
			StartDate:       v.Startdate,
			EndDate:         v.Enddate,
			AvgWaveSurvived: fmt.Sprintf("%.2f", v.Avgwave),
			HighestWave:     strconv.Itoa(int(v.Maxwave)),
			Count:           strconv.Itoa(int(v.Numoftestevents)),
		}
	}
	return &s
}
