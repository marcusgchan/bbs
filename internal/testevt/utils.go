package testevt

import (
	"fmt"
	"strconv"

	database "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal/testevt/views"
)

func TransformToTestEvtProps(data []database.GetTestEvtsRow) []testevt.TestEvtProps {
	length := len(data)
	mappedData := make([]testevt.TestEvtProps, length)
	for i, d := range data {
		mappedData[i] = testevt.TestEvtProps{
			ID:             d.ID,
			Environment:    d.Environment,
			Difficulty:     d.Difficulty,
			MainPlayerName: d.Mainplayer,
			StartedAt:      d.Startedat.String(),
			HasEnded:       d.Testresultid.Valid,
		}
	}
	return mappedData
}

func TransformToEvtResProps(evtData database.GetTestEvtResultsRow, playerData []database.GetTestEvtPlayerResultsRow) (testevt.TestEvtResProps, testevt.TestEvtTemplateRes, []testevt.TestEvtPlayerRes) {
	playerInfo := make([]testevt.TestEvtPlayerRes, len(playerData))
	for i, p := range playerData {
		playerInfo[i] = testevt.TestEvtPlayerRes{
			ID:       p.Player.ID,
			Name:     p.Player.Name,
			DiedTo:   p.PlayerTestResult.Diedto,
			WaveDied: strconv.Itoa(int(p.PlayerTestResult.Wavedied)),
		}
	}
	duration := evtData.TestResult.Endedat.Sub(evtData.TestEvent.Startedat)
	testResInfo := testevt.TestEvtResProps{
		TestEvtID:   evtData.TestEvent.ID,
		Difficulty:  evtData.TestEvent.Difficulty,
		Environment: evtData.TestEvent.Environment,
		StartedAt:   evtData.TestEvent.Startedat.String(),
		EndedAt:     evtData.TestResult.Endedat.String(),
		Duration:    fmt.Sprintf("%.2f", duration.Minutes()),
	}
	template := testevt.TestEvtTemplateRes{
		ID:    evtData.Template.ID,
		Owner: evtData.Template.Name,
		Data:  evtData.Template.Data,
	}
	return testResInfo, template, playerInfo
}
