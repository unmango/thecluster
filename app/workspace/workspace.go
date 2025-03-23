package workspace

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
	"github.com/unmango/thecluster/project"
)

type Model struct {
	ctx      context.Context
	err      error
	work     project.Workspace
	pwork    auto.Workspace
	settings *workspace.Project
}

func New(ctx context.Context, work project.Workspace) Model {
	return Model{
		ctx:  ctx,
		work: work,
	}
}

func (m Model) Init() tea.Cmd {
	return m.load
}

type (
	loaded   auto.Workspace
	settings *workspace.Project
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loaded:
		m.pwork = msg
		return m, m.psettings
	case settings:
		m.settings = msg
	case error:
		m.err = msg
	}

	return m, nil
}

func (m Model) View() string {
	if m.err != nil {
		return fmt.Sprintln(m.err)
	}
	if m.pwork == nil || m.settings == nil {
		return "Loading..."
	}

	return fmt.Sprint(m.settings.Name)
}

func (m Model) load() tea.Msg {
	work, err := m.work.Load(m.ctx)
	if err != nil {
		return err
	}

	return loaded(work)
}

func (m Model) psettings() tea.Msg {
	p, err := m.pwork.ProjectSettings(m.ctx)
	if err != nil {
		return err
	}

	return settings(p)
}
