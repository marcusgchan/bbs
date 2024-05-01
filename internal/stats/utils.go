package stats

import (
	"fmt"
	"strconv"

	database "github.com/marcusgchan/bbs/database/gen"
	stats "github.com/marcusgchan/bbs/internal/stats/views"
)

func TransformToVersionsField(data *[]database.Version) *[]stats.Version {
	versions := make([]stats.Version, len(*data))
	for i, v := range *data {
		versions[i] = stats.Version{
			Version:   v.Value,
			CreatedAt: v.Createdat.Time.String(),
		}
	}
	return &versions
}

func TransformToSingleField(data *database.GetStatsByVersionRow) *stats.Stats {
	if *data == (database.GetStatsByVersionRow{}) {
		return nil
	}
	return &stats.Stats{
		Version:         data.Version,
		StartDate:       data.Startdate,
		EndDate:         data.Enddate,
		AvgWaveSurvived: fmt.Sprintf("%.2f", data.Avgwave),
		HighestWave:     strconv.Itoa(int(data.Maxwave)),
		Count:           strconv.Itoa(int(data.Numoftestevents)),
	}
}

func TransformToMultiField(data *[]database.GetMostRecentStatsRow) *[]stats.Stats {
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

func TransformToStatsField(data *database.GetStatsByVersionRow) *stats.Stats {
	if data == nil {
		return nil
	}
	return &stats.Stats{
		Version:         data.Version,
		StartDate:       data.Startdate,
		EndDate:         data.Enddate,
		AvgWaveSurvived: fmt.Sprintf("%.2f", data.Avgwave),
		HighestWave:     strconv.Itoa(int(data.Maxwave)),
		Count:           strconv.Itoa(int(data.Numoftestevents)),
	}
}
