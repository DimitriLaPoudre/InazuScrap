package player

import (
	"github.com/DimitriLaPoudre/InazuScrap/internal/model"
	"github.com/charmbracelet/bubbles/table"
)

func PlayersToRows(players []model.Player) []table.Row {
	var rows []table.Row
	for _, player := range players {
		rows = append(rows, player.Row())
	}
	return rows
}
