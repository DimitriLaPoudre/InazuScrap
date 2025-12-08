package player

import (
	"slices"

	"github.com/DimitriLaPoudre/InazuScrap/internal/model"
)

func Sort(players []model.Player, stat string, asc bool) []model.Player {
	if asc {
		switch stat {
		case "name":
			slices.SortFunc(players, func(a, b model.Player) int {
				if a.Name < b.Name {
					return -1
				} else if a.Name == b.Name {
					return 0
				} else {

					return 1
				}
			})
		case "position":
			slices.SortFunc(players, func(a, b model.Player) int {
				return int(a.GetStat(stat)) - int(b.GetStat(stat))
			})
		case "element":
			slices.SortFunc(players, func(a, b model.Player) int {
				return int(a.GetStat(stat)) - int(b.GetStat(stat))
			})
		default:
			slices.SortFunc(players, func(a, b model.Player) int {
				return int(a.GetStat(stat)) - int(b.GetStat(stat))
			})
		}
	} else {
		slices.SortFunc(players, func(a, b model.Player) int {
			return int(b.GetStat(stat)) - int(a.GetStat(stat))
		})
	}
	return players
}
