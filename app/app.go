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
	filepicker filepicker.Model
	Proj       *project.Project
	err        error
}

func New() Model {
	return Model{
		filepicker: filepicker.New(),
	}
}

// Init implements tea.Model.
func (m Model) Init() tea.Cmd {
	if m.Proj != nil {
		return tea.Quit
	}

	return tea.Batch(
		load(context.Background()),
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

var header = lipgloss.NewStyle().
	Padding(1, 2)

// View implements tea.Model.
func (m Model) View() string {
	if m.Proj == nil {
		m.err = fmt.Errorf("no Project")
	}
	if m.err != nil {
		return m.err.Error()
	}

	var s strings.Builder
	s.WriteString(header.Render(
		fmt.Sprint("Project: ", m.Proj.Dir),
	))
	s.WriteString("\n")
	s.WriteString(m.filepicker.View())

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
