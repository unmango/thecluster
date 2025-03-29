package gtea

import tea "github.com/charmbracelet/bubbletea"

type finalModel tea.Model

func HaveFinalModel(expected tea.Model) finalModel {
	return expected
}

type finalOutput interface{}

func HaveFinalOutput(expected interface{}) finalOutput {
	return expected
}
