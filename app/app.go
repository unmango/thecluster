package app

import (
	"context"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/unmango/thecluster/app/selector"
	"github.com/unmango/thecluster/app/workspace"
	"github.com/unmango/thecluster/project"
)

type Model struct {
	selector selector.Model
	err      error

	Proj *project.Project
}

func New() Model {
	return Model{
		selector: selector.New(),
	}
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
	case loaded:
		m.Proj = msg
		return m, m.readDir
	case error:
		m.err = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.selector, cmd = m.selector.Update(msg)
	return m, cmd
}

var (
	container = lipgloss.NewStyle()
)

// View implements tea.Model.
func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintln(m.err)
	}
	if m.Proj == nil {
		return "no Project"
	}

	var s strings.Builder
	s.WriteString(m.Proj.Dir.Path())
	s.WriteString("\n")
	s.WriteString(m.selector.View())

	return container.Render(s.String())
}

type loaded *project.Project

func load(ctx context.Context) tea.Cmd {
	return func() tea.Msg {
		if proj, err := project.Load(ctx); err != nil {
			return err
		} else {
			return loaded(proj)
		}
	}
}

func (m Model) readDir() tea.Msg {
	ws, err := m.Proj.Workspaces()
	if err != nil {
		return err
	}

	ctx := context.Background()
	items := []tea.Model{}
	for w := range ws {
		items = append(items,
			workspace.New(ctx, w),
		)
	}

	return selector.Items(items)
}
