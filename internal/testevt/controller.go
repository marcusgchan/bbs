package testevt

import (
	"encoding/json"

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

type TestEvtReq struct {
	EnvironmentName string `json:"environmentName"`
	Template        struct {
		PlayerID string `json:"playerId"`
		Name     string `json:"name"`
		Data     string `json:"data"`
	}
	Date string `json:"date"`
}

func (h TestEventHandler) InjestTestEvent(c echo.Context) error {
	data := new(TestEvtReq)
	err := json.NewDecoder(c.Request().Body).Decode(data)
	if err != nil {
		return err
	}
	return nil
}
