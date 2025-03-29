package workspace

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/charmbracelet/bubbles/v2/list"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
	"github.com/unmango/thecluster/project"
)

var (
	selectedStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#fc00de"))
)

type delegate struct{}

func (delegate) Height() int  { return 1 }
func (delegate) Spacing() int { return 0 }

func (delegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	if m.Index() == index {
		s := selectedStyle.Render(i.rel)
		fmt.Fprint(w, "\u221f "+s)
	} else {
		fmt.Fprint(w, "\u221f "+i.rel)
	}
}

func (delegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return func() tea.Msg { return nil }
}

type item struct {
	work project.Workspace
	rel  string
}

func (i item) FilterValue() string {
	return string(i.work)
}

func NewList(proj *project.Project, workspaces []project.Workspace) list.Model {
	var (
		items  []list.Item
		height int
		width  int
	)

	if len(workspaces) >= 10 {
		height = 10
	} else {
		height = len(workspaces) + 2
	}

	for _, x := range workspaces {
		rel, err := filepath.Rel(proj.Dir.Path(), x.Path())
		if err != nil {
			panic(err)
		}

		items = append(items, item{x, rel})
		width = max(width, len(x))
	}

	m := list.New(items, delegate{}, width, height)
	m.SetShowFilter(false)
	m.SetShowHelp(false)
	m.SetShowStatusBar(false)
	m.SetShowTitle(false)

	return m
}
