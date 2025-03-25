package workspace

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ItemDelegate struct{}

func (ItemDelegate) Height() int  { return 1 }
func (ItemDelegate) Spacing() int { return 0 }

func (ItemDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	i, ok := item.(Model)
	if !ok {
		return
	}

	fmt.Fprint(w, i.View())
}

type tagged struct {
	index int
	msg   tea.Msg
}

func (ItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	switch msg := msg.(type) {
	case tagged:
		i := m.Items()[msg.index].(Model)
		n, cmd := i.Update(msg.msg)
		m.SetItem(msg.index, n.(Model))
		return tag(msg.index, cmd)
	}

	return nil
}

func (m Model) FilterValue() string {
	return m.work.String()
}

func tag(i int, cmd tea.Cmd) tea.Cmd {
	return func() tea.Msg {
		return tagged{i, cmd()}
	}
}
