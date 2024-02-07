package testevt

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal"
	"github.com/marcusgchan/bbs/internal/testevt/views"
)

type TestEventHandler struct {
	Q  *database.Queries
	DB *sql.DB
}

func (h TestEventHandler) ShowTestEvent(c echo.Context) error {
	return internal.Render(testevt.Page(), c)
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
		Createdat:   date,
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
	Date     string `json:"date"`
}

func (h TestEventHandler) CreateTemplate(c echo.Context) error {
	data := new(CreateTemplateReq)
	err := json.NewDecoder(c.Request().Body).Decode(data)
	if err != nil {
		return err
	}
	date, err := time.Parse(time.UnixDate, data.Date)
	if err != nil {
		return err
	}
	err = h.Q.CreatePlayerTemp(c.Request().Context(), database.CreatePlayerTempParams{
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

type CreateTestResultReq struct {
	TestResultID int64 `json:"testResultId"`
	MoneyEarned  int64 `json:"moneyEarned"`
	Damage       int64 `json:"damage"`
	Player       []struct {
		ID       string `json:"id"`
		WaveDied int64  `json:"waveDied"`
		DiedTo   string `json:"diedTo"`
	}
}

func (h TestEventHandler) CreatePlayerTestResult(c echo.Context) error {
	data := new(CreateTestResultReq)
	// Create general test result
	err := h.Q.CreateTestResult(c.Request().Context(), database.CreateTestResultParams{
		Moneyearned: data.MoneyEarned,
	})
	// Create player test result

	return nil
}

func (h TestEventHandler) CreateTestResult(c echo.Context) error {
	return nil
}
