package testevt

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	database "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal"
	views "github.com/marcusgchan/bbs/internal/testevt/views"
)

type TestEventHandler struct {
	Q  *database.Queries
	DB *sql.DB
}

func (h TestEventHandler) TestEvtPage(c echo.Context) error {
	const pageSize = 20
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	if internal.FromHTMX(c) {
		if infinite := c.QueryParam("infinite"); infinite != "true" {
			limit := pageSize
			offset := 0
			data, err := h.Q.GetTestEvts(c.Request().Context(), database.GetTestEvtsParams{
				Limit:  int64(limit),
				Offset: int64(offset),
			})
			if err != nil {
				return err
			}
			return internal.Render(views.TestEvtContent(TransformToTestEvtProps(data), page+1), c)
		}
		limit := pageSize
		offset := (page - 1) * limit
		data, err := h.Q.GetTestEvts(c.Request().Context(), database.GetTestEvtsParams{
			Limit:  int64(limit),
			Offset: int64(offset),
		})
		if err != nil {
			return err
		}
		return internal.Render(views.TestEvtRows(TransformToTestEvtProps(data), page+1), c)
	}

	limit := pageSize * page
	offset := 0
	data, err := h.Q.GetTestEvts(c.Request().Context(), database.GetTestEvtsParams{
		Limit:  int64(limit),
		Offset: int64(offset),
	})
	if err != nil {
		return err
	}
	return internal.Render(views.TestEvtPage(TransformToTestEvtProps(data), page+1), c)
}

func (h TestEventHandler) TestEvtResPage(c echo.Context) error {
	testEvtId := c.Param("testEventId")
	if len(testEvtId) == 0 {
		log.Printf("Missing query params testEventId")
		return c.String(500, "")
	}

	evtData, err := h.Q.GetTestEvtResults(c.Request().Context(), testEvtId)
	if err == sql.ErrNoRows {
		if internal.FromHTMX(c) {
			return internal.Render(views.OnGoingContent(), c)
		}
		return internal.Render(views.OnGoingPage(), c)
	}
	if err != nil {
		return err
	}

	playerData, err := h.Q.GetTestEvtPlayerResults(c.Request().Context(), evtData.TestResult.ID)
	if err != nil {
		return err
	}

	transEvtData, template, transPlayerData := TransformToEvtResProps(evtData, playerData)
	if internal.FromHTMX(c) {
		return internal.Render(views.TestEvtResContent(transEvtData, template, transPlayerData), c)
	}
	return internal.Render(views.TestEvtResPage(transEvtData, template, transPlayerData), c)
}

type CreateTestEvtReq struct {
	ID              string   `json:"id"`
	MainPlayerID    string   `json:"mainPlayerId"`
	PlayerIds       []string `json:"playerIds"`
	EnvironmentName string   `json:"environmentName"`
	TemplateID      string   `json:"templateId"`
	Difficulty      string   `json:"difficulty"`
	Version         string   `json:"version"`
	Date            int64    `json:"date"`
}

func (h TestEventHandler) CreateTestEvent(c echo.Context) error {
	data := new(CreateTestEvtReq)
	err := json.NewDecoder(c.Request().Body).Decode(data)
	if err != nil {
		return err
	}
	tx, err := h.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := h.Q.WithTx(tx)
	err = qtx.CreateVersion(c.Request().Context(), data.Version)
	if err != nil {
		return err
	}
	err = qtx.CreateTestEvt(c.Request().Context(), database.CreateTestEvtParams{
		ID:          data.ID,
		Templateid:  data.TemplateID,
		Environment: data.EnvironmentName,
		Startedat:   time.Unix(data.Date, 0),
		Version:     data.Version,
		Difficulty:  data.Difficulty,
	})
	if err != nil {
		return err
	}
	for _, playerId := range data.PlayerIds {
		err = qtx.CreatePlayerTestEvt(c.Request().Context(), database.CreatePlayerTestEvtParams{
			Playerid:    playerId,
			Testeventid: data.ID,
		})
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return c.NoContent(204)
}

type player struct {
	ID            string `json:"id"`
	WavesSurvived int64  `json:"wavesSurvived"`
	DiedTo        string `json:"diedTo"`
}

type CreateTestResultReq struct {
	TestEvtID   string   `json:"testEventId"`
	MoneyEarned int64    `json:"moneyEarned"`
	Players     []player `json:"players"`
	Date        int64    `json:"date"`
}

func (h TestEventHandler) CreatePlayerTestResult(c echo.Context) error {
	data := new(CreateTestResultReq)
	err := json.NewDecoder(c.Request().Body).Decode(data)
	if err != nil {
		return err
	}
	tx, err := h.DB.Begin()
	qtx := h.Q.WithTx(tx)
	if err != nil {
		return err
	}
	fmt.Printf("data %v", data)
	defer tx.Rollback()
	// Create general test result
	createdTestResId, err := qtx.CreateTestResult(c.Request().Context(), database.CreateTestResultParams{
		Moneyearned: data.MoneyEarned,
		Endedat:     time.Unix(data.Date, 0),
	})
	if err != nil {
		return err
	}
	id, err := qtx.UpdateTestEvtWithTestRes(c.Request().Context(), database.UpdateTestEvtWithTestResParams{
		Testresultid: sql.NullInt64{Int64: createdTestResId, Valid: true},
		ID:           data.TestEvtID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("bad %v %v | %v", id, data.TestEvtID, data)
		}
		return err
	}
	if id != data.TestEvtID {
		return fmt.Errorf("bad %v %v | %v", id, data.TestEvtID, data)
	}
	for _, p := range data.Players {
		// Create player test result
		err = qtx.CreatePlayerTestResult(c.Request().Context(), database.CreatePlayerTestResultParams{
			Playerid:      p.ID,
			Testresultid:  createdTestResId,
			Wavessurvived: p.WavesSurvived,
			Diedto:        p.DiedTo,
		})
		if err != nil {
			return err
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return c.NoContent(204)
}
