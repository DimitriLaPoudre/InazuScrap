package player

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
