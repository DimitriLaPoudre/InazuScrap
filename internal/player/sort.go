package player

import "slices"

func Sort(players []Player, stat string, asc bool) []Player {
	if asc {
		slices.SortFunc(players, func(a, b Player) int {
			return int(a.GetStat(stat)) - int(b.GetStat(stat))
		})
	} else {
		slices.SortFunc(players, func(a, b Player) int {
			return int(b.GetStat(stat)) - int(a.GetStat(stat))
		})
	}
	return players
}
