package player

import "slices"

func Score(p Player, weights map[PlayerStat]uint) int64 {
	var score int64
	for stat, w := range weights {
		score += int64(p.GetStat(stat)) * int64(w)
	}
	return score
}

func Sort(players []Player, weights map[PlayerStat]uint) []Player {
	scored := make([]Player, len(players))
	copy(scored, players)
	slices.SortFunc(scored, func(a, b Player) int {
		sa := Score(a, weights)
		sb := Score(b, weights)
		if sa > sb {
			return -1
		} else if sa < sb {
			return 1
		}
		if a.Name < b.Name {
			return -1
		} else if a.Name > b.Name {
			return 1
		}
		return 0
	})
	return scored
}
