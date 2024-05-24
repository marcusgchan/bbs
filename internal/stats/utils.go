package stats

import (
	"fmt"
	"strconv"

	database "github.com/marcusgchan/bbs/database/gen"
	stats "github.com/marcusgchan/bbs/internal/stats/views"
)

func createStats(data *database.GetTestEventsStatsRow) *stats.Stats {
	return &stats.Stats{
		Version:         data.Version,
		StartDate:       data.Startdate,
		EndDate:         data.Enddate,
		Count:           strconv.Itoa(int(data.Numoftestevents)),
		HighestWave:     strconv.Itoa(int(data.Maxwave)),
		AvgWaveSurvived: fmt.Sprintf("%.2f", data.Avgwave),
	}
}

func mergeStats(normal *[]database.GetTestEventsStatsRow, hard *[]database.GetTestEventsStatsRow) *[]stats.StatsByVersion {
	statsByVersion := make([]stats.StatsByVersion, 0)
	l := 0
	r := 0
	for l < len(*normal) || r < len(*hard) {
		if l < len(*normal) && r < len(*hard) {
			if (*normal)[l].Version == (*hard)[r].Version {
				statsByVersion = append(statsByVersion, stats.StatsByVersion{
					Version: (*normal)[l].Version,
					Normal:  createStats(&(*normal)[l]),
					Hard:    createStats(&(*hard)[r]),
				})
				l++
				r++
			} else if (*normal)[l].Version < (*hard)[r].Version {
				statsByVersion = append(statsByVersion, stats.StatsByVersion{
					Version: (*normal)[l].Version,
					Normal:  createStats(&(*normal)[l]),
				})
				l++
			} else {
				statsByVersion = append(statsByVersion, stats.StatsByVersion{
					Version: (*hard)[r].Version,
					Hard:    createStats(&(*hard)[r]),
				})
				r++
			}
		}

		if l < len(*normal) {
			statsByVersion = append(statsByVersion, stats.StatsByVersion{
				Version: (*normal)[l].Version,
				Normal:  createStats(&(*normal)[l]),
			})
			l++
		} else {
			statsByVersion = append(statsByVersion, stats.StatsByVersion{
				Version: (*hard)[r].Version,
				Hard:    createStats(&(*hard)[r]),
			})
			r++
		}
	}

	return &statsByVersion
}

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

func TransformToSingleField(data *database.GetTestEventStatsByVersionRow) *stats.Stats {
	if *data == (database.GetTestEventStatsByVersionRow{}) {
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

func TransformToMultiField(data *[]database.GetTestEventsStatsRow) *[]stats.Stats {
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

func TransformToStatsField(data *database.GetTestEventStatsByVersionRow) *stats.Stats {
	if *data == (database.GetTestEventStatsByVersionRow{}) {
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

func TransformToCatastropheField(data *[]database.GetCatastropheKillsRow) *[]stats.CatastropheDeaths {
	res := []stats.CatastropheDeaths{}
	prevVer := ""
	var cur *stats.CatastropheDeaths

	for _, row := range *data {
		if prevVer != row.Version {
			if cur != nil {
				res = append(res, *cur)
			}

			cur = &stats.CatastropheDeaths{
				Version:       row.Version,
				Catastrophies: []string{},
				Deaths:        []int{},
				TotalDeaths:   0,
			}
			prevVer = row.Version

			cur.Catastrophies = append(cur.Catastrophies, row.Catastrophe)
			cur.Deaths = append(cur.Deaths, int(row.Deaths))
			cur.TotalDeaths += int(row.Deaths)
			continue
		}

		cur.Catastrophies = append(cur.Catastrophies, row.Catastrophe)
		cur.Deaths = append(cur.Deaths, int(row.Deaths))
		cur.TotalDeaths += int(row.Deaths)
	}

	if cur != nil {
		res = append(res, *cur)
	}

	return &res
}

func parseCatastropheQueryParams(params string) int {
	val, err := strconv.Atoi(params)
	if err != nil || val < 1 || val > 10 {
		val = 3
	}
	return val
}

func parseTestEventQueryParams(params string) int {
	val, err := strconv.Atoi(params)
	if err != nil || val < 3 || val > 10 {
		val = 3
	}
	return val
}
