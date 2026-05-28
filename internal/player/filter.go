package player

import "slices"

type PlayerFilter struct {
	Positions []PlayerPosition
}

func Filter(l []Player, f PlayerFilter) []Player {
	r := []Player{}
	for _, p := range l {
		if len(f.Positions) != 0 && !slices.ContainsFunc(f.Positions, func(e PlayerPosition) bool {
			return e == p.Position
		}) {
			continue
		}
		r = append(r, p)
	}
	return r
}

func (f *PlayerFilter) TogglePosition(p PlayerPosition) {
	if i := slices.Index(f.Positions, p); i != -1 {
		f.Positions = slices.Delete(f.Positions, i, i+1)
	} else {
		f.Positions = append(f.Positions, p)
	}
}
