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

type core struct {
	table   table.Model
	players []player.Player
	filter  player.PlayerFilter
	sort    string
}

func (m core) Init() tea.Cmd {
	return nil
}

func (m core) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "n", "s", "k", "c", "t", "p", "shift+p", "a", "i":
			table := map[string]string{
				"n": "name",
				"s": "sum",
				"k": "kick",
				"c": "control",
				"t": "technique",
				"p": "pressure",
				"P": "physical",
				"a": "agility",
				"i": "intelligence",
			}
			m.sort = table[key]
			m.RefreshTable()
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

func (m *core) RefreshTable() {
	list := player.Filter(m.players, m.filter)
	rows := player.PlayersToRows(player.Sort(list, m.sort, false))
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

	columns := []table.Column{
		{Title: "Name", Width: 18},
		{Title: "Role", Width: 4},
		{Title: "Type", Width: 4},
		{Title: "Sum", Width: 4},
		{Title: "Kick", Width: 4},
		{Title: "Control", Width: 4},
		{Title: "Technique", Width: 4},
		{Title: "Pressure", Width: 4},
		{Title: "Physical", Width: 4},
		{Title: "Agility", Width: 4},
		{Title: "Intelligence", Width: 4},
	}

	players, err := player.Load("./players.json")
	if err != nil {
		fmt.Println(err)
		players = player.Scrap(8)
		err = player.Save("./players.json", players)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(len(players))

	rows := player.PlayersToRows(players)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
		table.WithStyles(s),
	)

	filter := player.PlayerFilter{Positions: []player.PlayerPosition{}}

	m := core{t, players, filter, "sum"}
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
