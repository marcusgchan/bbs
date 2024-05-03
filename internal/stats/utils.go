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

func TransformToCatastropheField(data *[]database.GetCatastropheKillsRow) *[]stats.CatastropheDeaths {
	res := []stats.CatastropheDeaths{}
	prevVer := ""
	var cur *stats.CatastropheDeaths

	for _, row := range *data {
		fmt.Printf("ver: %v\n", row.Version)
		fmt.Printf("prev ver |%v|\n", prevVer)
		if prevVer != row.Version {
			fmt.Printf("ver not the same. cur |%v|\n", cur)
			if cur != nil {
				fmt.Printf("appending prev content to array %v\n", *cur)
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
			fmt.Printf("data in loop: %v\n", res)
			continue
		}

		cur.Catastrophies = append(cur.Catastrophies, row.Catastrophe)
		cur.Deaths = append(cur.Deaths, int(row.Deaths))
		cur.TotalDeaths += int(row.Deaths)
	}

	if cur != nil {
		fmt.Printf("appending prev content to array %v\n", *cur)
		res = append(res, *cur)
	}
	fmt.Printf("data out loop: %v\n", res)

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
