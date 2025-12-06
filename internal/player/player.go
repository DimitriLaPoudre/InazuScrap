package player

import "fmt"

type Player struct {
	Name         string
	Position     string
	Element      string
	Kick         uint
	Control      uint
	Technique    uint
	Pressure     uint
	Physical     uint
	Agility      uint
	Intelligence uint
	Sum          uint
}

func (p Player) String() string {
	return fmt.Sprint(p.Name,
		"\n", p.Position,
		"\n", p.Element,
		"\ntotal: ", p.Sum,
		"\n\tkick: ", p.Kick,
		"\n\tcontrol: ", p.Control,
		"\n\ttechnique: ", p.Technique,
		"\n\tpressure: ", p.Pressure,
		"\n\tPhysical: ", p.Physical,
		"\n\tAgility: ", p.Agility,
		"\n\tIntelligence: ", p.Intelligence)
}
