package model

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

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

func (p *Player) Row() table.Row {
	return table.Row{p.Name, p.Position, p.Element,
		strconv.FormatUint(uint64(p.Kick), 10),
		strconv.FormatUint(uint64(p.Control), 10),
		strconv.FormatUint(uint64(p.Technique), 10),
		strconv.FormatUint(uint64(p.Pressure), 10),
		strconv.FormatUint(uint64(p.Physical), 10),
		strconv.FormatUint(uint64(p.Agility), 10),
		strconv.FormatUint(uint64(p.Intelligence), 10),
	}
}

func (p *Player) GetStat(stat string) uint {
	switch stat {
	case "sum":
		return p.Sum
	case "kick":
		return p.Kick
	case "control":
		return p.Control
	case "technique":
		return p.Technique
	case "pressure":
		return p.Pressure
	case "physical":
		return p.Physical
	case "agility":
		return p.Agility
	case "intelligence":
		return p.Intelligence
	default:
		return 0
	}
}
