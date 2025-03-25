package app

import (
	"context"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/unmango/thecluster/app/workspace"
	"github.com/unmango/thecluster/project"
)

type Model struct {
	ws  list.Model
	err error

	Proj *project.Project
}

func New() Model {
	return Model{}
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	if m.Proj == nil {
		return load(context.Background())
	}

	return nil
}

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case projectLoaded:
		m.Proj = msg
		return m, m.readDir
	case itemsLoaded:
		m.ws = list.New(msg, workspace.ItemDelegate{}, 20, 14)
	case error:
		m.err = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	return m, cmd
}

var (
	container = lipgloss.NewStyle()
	errorMsg  = lipgloss.NewStyle().Background(lipgloss.Color("#FF0000"))
)

// View implements tea.Model.
func (m Model) View() string {
	if m.Proj == nil {
		return "no Project"
	}

	var s strings.Builder
	s.WriteString(m.Proj.Dir.Path())
	s.WriteString("\n")
	s.WriteString(m.ws.View())

	if m.err != nil {
		s.WriteString("\n" + errorMsg.Render(m.err.Error()))
	}

	return container.Render(s.String())
}

type (
	projectLoaded *project.Project
	itemsLoaded   []list.Item
)

func load(ctx context.Context) tea.Cmd {
	return func() tea.Msg {
		if proj, err := project.Load(ctx); err != nil {
			return err
		} else {
			return projectLoaded(proj)
		}
	}
}

func (m Model) readDir() tea.Msg {
	ws, err := m.Proj.Workspaces()
	if err != nil {
		return err
	}

	ctx := context.Background()
	items := []list.Item{}
	for w := range ws {
		items = append(items,
			workspace.New(ctx, w),
		)
	}

	return itemsLoaded(items)
}
