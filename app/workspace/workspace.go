package workspace

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/unmango/thecluster/project"
)

type Model struct {
	ctx   context.Context
	err   error
	work  project.Workspace
	pwork auto.Workspace
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

type loaded auto.Workspace

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loaded:
		m.pwork = msg
	case error:
		m.err = msg
	}

	return m, nil
}

func (m Model) View() string {
	if m.pwork == nil {
		return "Loading..."
	}

	return m.pwork.WorkDir()
}

func (m Model) load() tea.Msg {
	work, err := m.work.Load(m.ctx)
	if err != nil {
		return err
	}

	return loaded(work)
}
