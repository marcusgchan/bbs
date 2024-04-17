package template

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
	database "github.com/marcusgchan/bbs/database/gen"
)

type TemplateHandler struct {
	Q  *database.Queries
	DB *sql.DB
}

type CreateTemplateReq struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	PlayerID string `json:"playerId"`
	Date     int64  `json:"date"`
}

func (h *TemplateHandler) Create(c echo.Context) error {
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
