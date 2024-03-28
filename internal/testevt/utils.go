package testevt

import (
	database "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal/testevt/views"
)

func TransformToTestEvtProps(data []database.TestEvent) []testevt.TestEvtProps {
	length := len(data)
	mappedData := make([]testevt.TestEvtProps, length)
	for i, d := range data {
		mappedData[i] = testevt.TestEvtProps{
			ID:          d.ID,
			Environment: d.Environment,
			Difficulty:  d.Difficulty,
			StartedAt:   d.Startedat.String(),
			HasEnded:    d.Testresultid.Valid,
		}
	}
	return mappedData
}
