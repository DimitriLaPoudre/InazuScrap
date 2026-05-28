package player

import (
	"github.com/charmbracelet/bubbles/table"
)

func PlayersToRows(players []Player) []table.Row {
	var rows []table.Row
	for _, player := range players {
		rows = append(rows, player.Row())
	}
	return rows
}
