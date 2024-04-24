package player

import (
	database "github.com/marcusgchan/bbs/database/gen"
	"github.com/marcusgchan/bbs/internal/player/views"
)

func TransformToPlayerProps(data []database.Player) []player.PlayerProps {
	players := make([]player.PlayerProps, len(data))
	for i, p := range data {
		players[i] = player.PlayerProps{
			ID:   p.ID,
			Name: p.Name,
		}
	}
	return players
}
