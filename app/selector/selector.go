package selector

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	item = lipgloss.NewStyle().
		Background(lipgloss.Color("#0f0f0f"))

	selected = lipgloss.NewStyle()
)

type Model struct {
	items    []tea.Model
	selected int
}

func New() Model {
	return Model{
		items:    []tea.Model{},
		selected: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

type Items []tea.Model

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case Items:
		m.items = msg
		for _, i := range m.items {
			cmds = append(cmds, i.Init())
		}
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

	for i, item := range m.items {
		var cmd tea.Cmd
		m.items[i], cmd = item.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	b := strings.Builder{}
	for idx, i := range m.items {
		row := fmt.Sprint("\u221f ", i.View())
		if idx == m.selected {
			b.WriteString(selected.Render(row + " <"))
		} else {
			b.WriteString(item.Render(row))
		}
		b.WriteString("\n")
	}

	return b.String()
}
