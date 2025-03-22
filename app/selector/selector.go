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
			Background(lipgloss.Color("#f0f0f0"))
)

type Model struct {
	items    []string
	selected int
}

func New() Model {
	return Model{
		items:    []string{},
		selected: 1,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

type Items []string

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case Items:
		m.items = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			m.selected -= 1
		case "down":
			m.selected += 1
		}
		m.selected = max(0, m.selected)
		m.selected = min(len(m.items)-1, m.selected)
	}

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
