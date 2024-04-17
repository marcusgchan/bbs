package testevt

import (
	"database/sql"
	"encoding/json"
	"log"
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

func (h TestEventHandler) GetTestEvtPage(c echo.Context) error {
	data, err := h.Q.GetTestEvts(c.Request().Context())
	if err != nil {
		return err
	}

	if internal.FromHTMX(c) {
		return internal.Render(views.TestEvtContent(TransformToTestEvtProps(data)), c)
	}

	return internal.Render(views.TestEvtPage(TransformToTestEvtProps(data)), c)
}

func (h TestEventHandler) GetTestEvtResPage(c echo.Context) error {
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
	Players         []string `json:"players"`
	EnvironmentName string   `json:"environmentName"`
	TemplateID      string   `json:"templateId"`
	Difficulty      string   `json:"difficulty"`
	Date            string   `json:"date"`
}

func (h TestEventHandler) CreateTestEvent(c echo.Context) error {
	data := new(CreateTestEvtReq)
	err := json.NewDecoder(c.Request().Body).Decode(data)
	if err != nil {
		return err
	}
	date, err := time.Parse(time.UnixDate, data.Date)
	if err != nil {
		return err
	}
	tx, err := h.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	qtx := h.Q.WithTx(tx)
	err = qtx.CreateTestEvt(c.Request().Context(), database.CreateTestEvtParams{
		Environment: data.EnvironmentName,
		Startedat:   date,
		Difficulty:  data.Difficulty,
	})
	if err != nil {
		return err
	}
	tx.Commit()
	return c.NoContent(204)
}

type CreateTemplateReq struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	PlayerID string `json:"playerId"`
	Date     int64  `json:"date"`
}

func (h TestEventHandler) CreateTemplate(c echo.Context) error {
	data := new(CreateTemplateReq)
	err := json.NewDecoder(c.Request().Body).Decode(data)
	if err != nil {
		return err
	}
	date := time.Unix(data.Date, 0)
	err = h.Q.CreatePlayerTemp(c.Request().Context(), database.CreatePlayerTempParams{
		ID:        data.ID,
		Playerid:  data.PlayerID,
		Data:      data.Data,
		Name:      data.Name,
		Createdat: date,
	})
	if err != nil {
		return err
	}
	return c.NoContent(204)
}

type player struct {
	ID       string `json:"id"`
	WaveDied int64  `json:"waveDied"`
	DiedTo   string `json:"diedTo"`
}

type CreateTestResultReq struct {
	TestEvtID   string   `json:"testEvtId"`
	MoneyEarned int64    `json:"moneyEarned"`
	Players     []player `json:"players"`
	Date        string   `json:"date"`
}

func (h TestEventHandler) CreatePlayerTestResult(c echo.Context) error {
	data := new(CreateTestResultReq)
	err := json.NewDecoder(c.Request().Body).Decode(data)
	if err != nil {
		return err
	}
	date, err := time.Parse(time.UnixDate, data.Date)
	if err != nil {
		return err
	}
	tx, err := h.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	// Create general test result
	createdTestResId, err := h.Q.CreateTestResult(c.Request().Context(), database.CreateTestResultParams{
		Moneyearned: data.MoneyEarned,
		Endedat:     date,
	})
	if err != nil {
		return err
	}
	err = h.Q.UpdateTestEvtWithTestRes(c.Request().Context(), database.UpdateTestEvtWithTestResParams{
		Testresultid: sql.NullInt64{Int64: createdTestResId, Valid: true},
		ID:           data.TestEvtID,
	})
	for _, p := range data.Players {
		// Create player test result
		h.Q.CreatePlayerTestResult(c.Request().Context(), database.CreatePlayerTestResultParams{
			Playerid:     p.ID,
			Testresultid: createdTestResId,
			Wavedied:     p.WaveDied,
			Diedto:       p.DiedTo,
		})
		if err != nil {
			return err
		}
	}
	return c.NoContent(204)
}
