package player

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	database "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal"
	"github.com/marcusgchan/bbs/internal/player/views"
)

type PlayerHandler struct {
	Q  *database.Queries
	DB *sql.DB
}

func (h PlayerHandler) ShowPlayerList(c echo.Context) error {
	data, err := h.Q.GetPlayers(c.Request().Context())
	if err != nil {
		emptyData := []player.PlayerProps{}
		return internal.Render(player.PlayersTable(emptyData), c)
	}
	players := make([]player.PlayerProps, len(data))
	for i, d := range data {
		players[i] = player.PlayerProps{
			ID:   d.ID,
			Name: d.Name,
		}
	}
	return internal.Render(player.PlayersTable(players), c)
}

type CreatePlayerReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h PlayerHandler) CreatePlayer(c echo.Context) error {
	data := new(CreatePlayerReq)
	err := json.NewDecoder(c.Request().Body).Decode(data)
	if err != nil {
		return err
	}
	err = h.Q.CreatePlayer(c.Request().Context(), database.CreatePlayerParams{
		ID:   data.ID,
		Name: data.Name,
	})
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}
