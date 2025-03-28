package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/unmango/thecluster/app/selector"
)

type Model struct {
	selector selector.Model
}

func New() Model {
	return Model{
		selector: selector.New(),
	}
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	return m.selector.Init()
}

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	s, cmd := m.selector.Update(msg)
	m.selector = s.(selector.Model)
	return m, cmd
}

// View implements tea.Model.
func (m Model) View() string {
	return m.selector.View()
}
