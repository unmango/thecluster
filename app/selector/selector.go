package selector

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/unmango/thecluster/app/workspace"
	"github.com/unmango/thecluster/project"
)

var (
	container = lipgloss.NewStyle()
)

type Model struct {
	selector list.Model
	err      error

	Proj *project.Project
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	if m.Proj == nil {
		return load
	}

	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loaded:
		m.Proj = msg.proj
		m.selector = workspace.NewList(m.Proj, msg.ws)
	case error:
		m.err = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	if m.Proj != nil {
		m.selector, cmd = m.selector.Update(msg)
	}
	return m, cmd
}

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

type loaded struct {
	proj *project.Project
	ws   []project.Workspace
}

func load() tea.Msg {
	ctx := context.Background()
	proj, err := project.Load(ctx)
	if err != nil {
		return err
	}

	ws, err := proj.Workspaces()
	if err != nil {
		return err
	}

	return loaded{proj, slices.Collect(ws)}
}
