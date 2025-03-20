package header

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	header = lipgloss.NewStyle().
		Padding(1, 2)
)

type Model struct {
	Title string
}

func New(title string) Model {
	return Model{Title: title}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return header.Render(m.Title)
}
