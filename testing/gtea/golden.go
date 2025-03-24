package gtea

import (
	"fmt"
	"io"
	"reflect"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"
)

func RequireGolden(tm *teatest.TestModel) {
	ginkgo.GinkgoHelper()
	tb := ginkgo.GinkgoTB()

	out, err := io.ReadAll(tm.FinalOutput(tb))
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	teatest.RequireEqualOutput(tb, out)
}

type golden struct{}

func BeGolden() golden {
	return golden{}
}

// Match implements types.GomegaMatcher.
func (g golden) Match(actual interface{}) (success bool, err error) {
	tm, ok := actual.(*teatest.TestModel)
	if !ok {
		return false, fmt.Errorf("expected a *teatest.TestModel, got %v", reflect.TypeOf(actual))
	}

	panic("unimplemented")
}

// FailureMessage implements types.GomegaMatcher.
func (g golden) FailureMessage(actual interface{}) (message string) {
	panic("unimplemented")
}

// NegatedFailureMessage implements types.GomegaMatcher.
func (g golden) NegatedFailureMessage(actual interface{}) (message string) {
	panic("unimplemented")
}

type finalModel tea.Model

func HaveFinalModel(expected tea.Model) finalModel {
	return expected
}

type finalOutput interface{}

func HaveFinalOutput(expected interface{}) finalOutput {
	return expected
}
