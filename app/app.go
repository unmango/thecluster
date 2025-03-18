package app

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/unmango/thecluster/project"
)

type Model struct {
	Proj *project.Project
	err  error
}

// Init implements tea.Model.
func (m *Model) Init() tea.Cmd {
	if m.Proj != nil {
		return tea.Quit
	}

	return load(context.Background())
}

// Update implements tea.Model.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loaded:
		m.Proj = msg
		return m, tea.Quit
	case error:
		m.err = msg
		return m, tea.Quit
	}

	return m, nil
}

// View implements tea.Model.
func (m *Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	if m.Proj == nil {
		return ""
	}

	return fmt.Sprint("Project: ", m.Proj.Dir)
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
