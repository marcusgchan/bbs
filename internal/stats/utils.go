package stats

import (
	"fmt"
	"strconv"

	database "github.com/marcusgchan/bbs/database/gen"
	stats "github.com/marcusgchan/bbs/internal/stats/views"
)

func createStats(data *database.GetTestEventsStatsRow) *stats.GeneralStats {
	return &stats.GeneralStats{
		Version:         data.Version,
		StartDate:       data.Startdate,
		EndDate:         data.Enddate,
		Count:           strconv.Itoa(int(data.Numoftestevents)),
		HighestWave:     strconv.Itoa(int(data.Maxwave)),
		AvgWaveSurvived: fmt.Sprintf("%.2f", data.Avgwave),
	}
}

func mergeGeneralStats(normal *[]database.GetTestEventsStatsRow, hard *[]database.GetTestEventsStatsRow) *[]stats.StatsByVersion {
	statsByVersion := make([]stats.StatsByVersion, 0)
	l := 0
	r := 0
	for l < len(*normal) || r < len(*hard) {
		if l < len(*normal) && r < len(*hard) {
			if (*normal)[l].Version == (*hard)[r].Version {
				statsByVersion = append(statsByVersion, stats.StatsByVersion{
					Version: (*normal)[l].Version,
					Normal: &stats.Stats{
						GeneralStats: createStats(&((*normal)[l])),
					},
					Hard: &stats.Stats{
						GeneralStats: createStats(&((*hard)[r])),
					},
				})
				l++
				r++
			} else if (*normal)[l].Version < (*hard)[r].Version {
				statsByVersion = append(statsByVersion, stats.StatsByVersion{
					Version: (*normal)[l].Version,
					Normal: &stats.Stats{
						GeneralStats: createStats(&((*normal)[l])),
					},
				})
				l++
			} else {
				statsByVersion = append(statsByVersion, stats.StatsByVersion{
					Version: (*hard)[r].Version,
					Hard: &stats.Stats{
						GeneralStats: createStats(&((*hard)[r])),
					},
				})
				r++
			}
		} else if l < len(*normal) {
			statsByVersion = append(statsByVersion, stats.StatsByVersion{
				Version: (*normal)[l].Version,
				Normal: &stats.Stats{
					GeneralStats: createStats(&((*normal)[l])),
				},
			})
			l++
		} else {
			statsByVersion = append(statsByVersion, stats.StatsByVersion{
				Version: (*hard)[r].Version,
				Hard: &stats.Stats{
					GeneralStats: createStats(&((*hard)[r])),
				},
			})
			r++
		}
	}

	return &statsByVersion
}

type verDiffToCatKey struct {
	version    string
	difficulty string
}

func addCatastropheDeathsToStats(statsByVersion *[]stats.StatsByVersion, catastropheDeaths *[]database.GetCatastropheKillsRow) {
	verDiffToCat := map[verDiffToCatKey]*stats.CatastropheDeaths{}
	for _, catDeath := range *catastropheDeaths {
		key := verDiffToCatKey{
			version:    catDeath.Version,
			difficulty: catDeath.Difficulty,
		}
		cat := verDiffToCat[key]
		if cat == nil {
			verDiffToCat[key] = &stats.CatastropheDeaths{
				Deaths: []int{
					int(catDeath.Deaths),
				},
				Catastrophies: []string{
					catDeath.Catastrophe,
				},
				TotalDeaths: 1,
			}
		} else {
			cat.Catastrophies = append(cat.Catastrophies, catDeath.Catastrophe)
			cat.Deaths = append(cat.Deaths, int(catDeath.Deaths))
			cat.TotalDeaths++
		}
	}

	for _, s := range *statsByVersion {
		if s.Hard != nil {
			key := verDiffToCatKey{
				version:    s.Version,
				difficulty: "hard",
			}
			s.Hard.CatastropheDeaths = verDiffToCat[key]
		}
		if s.Normal != nil {
			key := verDiffToCatKey{
				version:    s.Version,
				difficulty: "normal",
			}
			s.Normal.CatastropheDeaths = verDiffToCat[key]
		}
	}
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

func parseTestEventQueryParams(params string) int {
	val, err := strconv.Atoi(params)
	if err != nil || val < 1 || val > 10 {
		val = 1
	}
	return val
}
