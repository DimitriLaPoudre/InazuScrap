package main

import (
	"fmt"

	"github.com/DimitriLaPoudre/InazuScrap/internal/player"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

var statKeys = map[string]player.PlayerStat{
	"1": player.StatSum,
	"2": player.StatKick,
	"3": player.StatControl,
	"4": player.StatTechnique,
	"5": player.StatPressure,
	"6": player.StatPhysical,
	"7": player.StatAgility,
	"8": player.StatIntelligence,
}

var statKeysShift = map[string]player.PlayerStat{
	"!": player.StatSum,
	"@": player.StatKick,
	"#": player.StatControl,
	"$": player.StatTechnique,
	"%": player.StatPressure,
	"^": player.StatPhysical,
	"&": player.StatAgility,
	"*": player.StatIntelligence,
}

type core struct {
	table    table.Model
	players  []player.Player
	filter   player.PlayerFilter
	sortStat map[player.PlayerStat]uint
}

func (m core) Init() tea.Cmd {
	return nil
}

func (m core) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "1", "2", "3", "4", "5", "6", "7", "8":
			stat := statKeys[key]
			if m.sortStat[stat] < 9 {
				m.sortStat[stat]++
				m.RefreshTable()
			}
		case "!", "@", "#", "$", "%", "^", "&", "*":
			stat := statKeysShift[key]
			if m.sortStat[stat] > 0 {
				m.sortStat[stat]--
				if m.sortStat[stat] == 0 {
					delete(m.sortStat, stat)
				}
				m.RefreshTable()
			}
		case "alt+f":
			m.filter.TogglePosition(player.FW)
			m.RefreshTable()
		case "alt+m":
			m.filter.TogglePosition(player.MF)
			m.RefreshTable()
		case "alt+d":
			m.filter.TogglePosition(player.DF)
			m.RefreshTable()
		case "alt+g":
			m.filter.TogglePosition(player.GK)
			m.RefreshTable()
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *core) columns() []table.Column {
	w := func(stat player.PlayerStat, key, name string, width int) string {
		return fmt.Sprintf("(%s)%-*s [%d]", key, width, name, m.sortStat[stat])
	}
	return []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Role", Width: 6},
		{Title: "Type", Width: 6},
		{Title: w(player.StatSum, "1", "Sum", 5), Width: 13},
		{Title: w(player.StatKick, "2", "Kick", 5), Width: 13},
		{Title: w(player.StatControl, "3", "Control", 8), Width: 16},
		{Title: w(player.StatTechnique, "4", "Technique", 10), Width: 19},
		{Title: w(player.StatPressure, "5", "Pressure", 9), Width: 18},
		{Title: w(player.StatPhysical, "6", "Physical", 9), Width: 18},
		{Title: w(player.StatAgility, "7", "Agility", 8), Width: 16},
		{Title: w(player.StatIntelligence, "8", "Intelligence", 12), Width: 23},
	}
}

func (m *core) RefreshTable() {
	m.table.SetColumns(m.columns())
	list := player.Filter(m.players, m.filter)
	rows := player.PlayersToRows(player.Sort(list, m.sortStat))
	m.table.SetRows(rows)
}

func (m core) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func main() {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	players, err := player.Load("./players.json")
	if err != nil {
		fmt.Println(err)
		players = player.Scrap(8)
		err = player.Save("./players.json", players)
		if err != nil {
			fmt.Println(err)
		}
	}
	players = player.Deduplicate(players)

	t := table.New(
		table.WithFocused(true),
		table.WithStyles(s),
	)

	m := core{
		table:    t,
		players:  players,
		filter:   player.PlayerFilter{Positions: []player.PlayerPosition{}},
		sortStat: map[player.PlayerStat]uint{},
	}
	m.RefreshTable()
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
