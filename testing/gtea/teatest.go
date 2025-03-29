package gtea

import tea "github.com/charmbracelet/bubbletea"

type finalModel struct {
	m tea.Model
}

func HaveFinalModel(expected tea.Model) finalModel {
	return finalModel{expected}
}

// Match implements types.GomegaMatcher.
func (g *finalModel) Match(actual interface{}) (success bool, err error) {
	panic("unimplemented")
}

// FailureMessage implements types.GomegaMatcher.
func (g *finalModel) FailureMessage(actual interface{}) (message string) {
	panic("unimplemented")
}

// NegatedFailureMessage implements types.GomegaMatcher.
func (g *finalModel) NegatedFailureMessage(actual interface{}) (message string) {
	panic("unimplemented")
}

type finalOutput struct {
	out []byte
}

func HaveFinalOutput(expected []byte) finalOutput {
	return finalOutput{expected}
}

// Match implements types.GomegaMatcher.
func (g *finalOutput) Match(actual interface{}) (success bool, err error) {
	panic("unimplemented")
}

// FailureMessage implements types.GomegaMatcher.
func (g *finalOutput) FailureMessage(actual interface{}) (message string) {
	panic("unimplemented")
}

// NegatedFailureMessage implements types.GomegaMatcher.
func (g *finalOutput) NegatedFailureMessage(actual interface{}) (message string) {
	panic("unimplemented")
}
