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
	DB *database.Queries
}

func (h TestEventHandler) ShowTestEvent(c echo.Context) error {
	return internal.Render(testevt.Page(), c)
}

type CreateTestEvtReq struct {
	ID              string `json:"id"`
	EnvironmentName string `json:"environmentName"`
	TemplateID      string `json:"templateId"`
	Difficulty      string `json:"difficulty"`
	Date            string `json:"date"`
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
	err = h.DB.CreateTestEvt(c.Request().Context(), database.CreateTestEvtParams{
		Environment: data.EnvironmentName,
		Createdat:   sql.NullTime{Time: date, Valid: true},
		Difficulty:  data.Difficulty,
	})
	if err != nil {
		return err
	}
	return c.NoContent(204)
}

type CreateTemplateReq struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	PlayerID string `json:"playerId"`
}

func (h TestEventHandler) CreateTemplate(c echo.Context) error {
	data := new(CreateTemplateReq)
	err := json.NewDecoder(c.Request().Body).Decode(data)
	if err != nil {
		return err
	}
	err = h.DB.CreatePlayerTemp(c.Request().Context(), database.CreatePlayerTempParams{
		Playerid:  data.PlayerID,
		Data:      data.Data,
		Name:      data.Name,
		Createdat: sql.NullTime{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return err
	}
	return c.NoContent(204)
}
