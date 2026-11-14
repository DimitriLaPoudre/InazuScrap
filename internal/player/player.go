package player

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
)

type PlayerPosition string

const (
	GK PlayerPosition = "GK"
	DF PlayerPosition = "DF"
	MF PlayerPosition = "MF"
	FW PlayerPosition = "FW"
)

type PlayerStat string

const (
	StatKick         PlayerStat = "kick"
	StatControl      PlayerStat = "control"
	StatTechnique    PlayerStat = "technique"
	StatPressure     PlayerStat = "pressure"
	StatPhysical     PlayerStat = "physical"
	StatAgility      PlayerStat = "agility"
	StatIntelligence PlayerStat = "intelligence"
	StatSum          PlayerStat = "sum"
)

type Player struct {
	Name         string
	Position     PlayerPosition
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
	return table.Row{p.Name, string(p.Position), p.Element,
		strconv.FormatUint(uint64(p.Sum), 10),
		strconv.FormatUint(uint64(p.Kick), 10),
		strconv.FormatUint(uint64(p.Control), 10),
		strconv.FormatUint(uint64(p.Technique), 10),
		strconv.FormatUint(uint64(p.Pressure), 10),
		strconv.FormatUint(uint64(p.Physical), 10),
		strconv.FormatUint(uint64(p.Agility), 10),
		strconv.FormatUint(uint64(p.Intelligence), 10),
	}
}

func (p *Player) GetStat(stat PlayerStat) uint {
	switch stat {
	case StatSum:
		return p.Sum
	case StatKick:
		return p.Kick
	case StatControl:
		return p.Control
	case StatTechnique:
		return p.Technique
	case StatPressure:
		return p.Pressure
	case StatPhysical:
		return p.Physical
	case StatAgility:
		return p.Agility
	case StatIntelligence:
		return p.Intelligence
	default:
		return 0
	}
}
