package app

import (
	"context"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/unmango/thecluster/app/header"
	"github.com/unmango/thecluster/app/selector"
	"github.com/unmango/thecluster/project"
)

type Model struct {
	ctx      context.Context
	header   header.Model
	selector selector.Model
	err      error

	Proj *project.Project
}

func New(ctx context.Context) Model {
	return Model{
		ctx:      ctx,
		header:   header.New(),
		selector: selector.New(),
	}
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	if m.Proj != nil {
		return tea.Quit
	}

	return tea.Batch(
		load(m.ctx),
	)
}

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loaded:
		m.Proj = msg
		m.header.Title = "Project: " + msg.Dir.Path()
		return m, m.readDir
	case error:
		m.err = msg
		return m, tea.Quit
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
		return m.err.Error()
	}
	if m.Proj == nil {
		return "no Project"
	}

	var s strings.Builder
	s.WriteString(m.header.View())
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

	items := []string{}
	for w := range ws {
		items = append(items, w)
	}

	return selector.Items(items)
}
