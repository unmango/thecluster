package selector

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	item = lipgloss.NewStyle().
		Background(lipgloss.Color("#0f0f0f"))

	selected = lipgloss.NewStyle().
			Background(lipgloss.Color("#000000"))
)

type Model struct {
	items    []string
	selected int
}

func New() Model {
	return Model{
		items:    []string{"TEST A", "TEST B"},
		selected: 1,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	b := strings.Builder{}
	for idx, i := range m.items {
		if idx == m.selected {
			b.WriteString(selected.Render(i))
		} else {
			b.WriteString(item.Render(i))
		}
		b.WriteString("\n")
	}

	return b.String()
}
