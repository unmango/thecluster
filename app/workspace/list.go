package workspace

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/unmango/thecluster/project"
)

type delegate struct{}

func (delegate) Height() int  { return 1 }
func (delegate) Spacing() int { return 0 }

func (delegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	if i, ok := listItem.(item); ok {
		fmt.Fprint(w, i)
	}
}

func (delegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

type item project.Workspace

func (i item) FilterValue() string {
	return string(i)
}

func NewList(workspaces []project.Workspace) list.Model {
	var (
		items  []list.Item
		height int = len(workspaces)
		width  int
	)

	for _, x := range workspaces {
		items = append(items, item(x))
		width = max(width, len(x))
	}

	m := list.New(items, delegate{}, width, height)
	m.SetShowFilter(false)
	m.SetShowHelp(false)
	m.SetShowStatusBar(false)
	m.SetShowTitle(false) // Might switch to this later

	return m
}
