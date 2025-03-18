package app

import tea "github.com/charmbracelet/bubbletea/v2"

type Model struct{}

// Init implements tea.Model.
func (m *Model) Init() tea.Cmd {
	return tea.Quit
}

// Update implements tea.Model.
func (m *Model) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

// View implements tea.Model.
func (m *Model) View() string {
	return "TEST"
}
