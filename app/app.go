package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/unmango/thecluster/project"
)

type Model struct {
	ctx        context.Context
	filepicker filepicker.Model
	err        error

	Proj *project.Project
}

func New(ctx context.Context) Model {
	return Model{
		ctx:        ctx,
		filepicker: filepicker.New(),
	}
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	if m.Proj != nil {
		return tea.Quit
	}

	return tea.Batch(
		load(m.ctx),
		m.filepicker.Init(),
	)
}

// Update implements tea.Model.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loaded:
		m.Proj = msg
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
	m.filepicker, cmd = m.filepicker.Update(msg)

	return m, cmd
}

var (
	header = lipgloss.NewStyle().
		Padding(1, 2)

	selected = lipgloss.NewStyle().
			Padding(0, 25).
			Margin(1, 0).
			Background(lipgloss.Color("#0f0f0f"))
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
	s.WriteString(header.Render(
		fmt.Sprint("Project: ", m.Proj.Dir),
	))
	s.WriteString("\n")
	s.WriteString(selected.Render("TEST"))

	return s.String()
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
