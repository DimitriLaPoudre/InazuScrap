package main

import (
	"fmt"

	"github.com/DimitriLaPoudre/InazuScrap/internal/app"
	"github.com/DimitriLaPoudre/InazuScrap/internal/model"
	"github.com/DimitriLaPoudre/InazuScrap/internal/player"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type Scene uint

const (
	SceneList   Scene = 0
	SceneSort   Scene = 1
	SceneFilter Scene = 2
)

type core struct {
	table   table.Model
	players []model.Player
	scene   Scene
}

func (m core) Init() tea.Cmd {
	return nil
}

func (m core) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
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

func (m core) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func main() {
	columns := []table.Column{
		{Title: "Name", Width: 18},
		{Title: "Role", Width: 4},
		{Title: "Type", Width: 4},
		{Title: "Kick", Width: 4},
		{Title: "Control", Width: 4},
		{Title: "Technique", Width: 4},
		{Title: "Pressure", Width: 4},
		{Title: "Physical", Width: 4},
		{Title: "Agility", Width: 4},
		{Title: "Intelligence", Width: 4},
	}

	fmt.Println("Scrap Start")
	players := app.Scrap()
	fmt.Println("Scrap End")

	players = player.Sort(players, "name", true)

	rows := player.PlayersToRows(players)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

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
	t.SetStyles(s)

	m := core{t, players, SceneList}
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
