package player

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	database "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal"
	"github.com/marcusgchan/bbs/internal/player/views"
)

type PlayerHandler struct {
	Q  *database.Queries
	DB *sql.DB
}

func (h PlayerHandler) PlayerListPage(c echo.Context) error {
	const pageSize = 20
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	if internal.FromHTMX(c) {
		if infinite := c.QueryParam("infinite"); infinite != "true" {
			limit := pageSize
			offset := 0
			data, err := h.Q.GetPlayers(c.Request().Context(), database.GetPlayersParams{
				Offset: int64(offset),
				Limit:  int64(limit),
			})
			if err != nil {
				return err
			}

			return internal.Render(player.PlayersContent(TransformToPlayerProps(data), page+1), c)
		}

		limit := pageSize
		offset := (page - 1) * pageSize
		data, err := h.Q.GetPlayers(c.Request().Context(), database.GetPlayersParams{
			Offset: int64(offset),
			Limit:  int64(limit),
		})
		if err != nil {
			return err
		}

		return internal.Render(player.PlayerRows(TransformToPlayerProps(data), page+1), c)
	}

	limit := pageSize
	offset := 0
	data, err := h.Q.GetPlayers(c.Request().Context(), database.GetPlayersParams{
		Offset: int64(offset),
		Limit:  int64(limit),
	})
	if err != nil {
		return err
	}
	return internal.Render(player.PlayersPage(TransformToPlayerProps(data), page+1), c)
}

func (h PlayerHandler) PlayerInfoPage(c echo.Context) error {
	playerId := c.Param("playerId")
	if len(playerId) == 0 {
		log.Printf("playerId not not found in params")
		return c.String(500, "")
	}
	return nil
}

func (h PlayerHandler) BlockBuyEvt(c echo.Context) error {
	return nil
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
