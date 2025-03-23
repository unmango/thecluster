package selector_test

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/thecluster/app/selector"
	"github.com/unmango/thecluster/testing/gtea"
)

type wrapper struct{ selector.Model }

func (w wrapper) Init() tea.Cmd { return w.Model.Init() }
func (w wrapper) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	w.Model, cmd = w.Model.Update(msg)
	return w, cmd
}
func (w wrapper) View() string { return w.Model.View() }

type item string

func (item) Init() tea.Cmd                         { return nil }
func (i item) Update(tea.Msg) (tea.Model, tea.Cmd) { return i, nil }
func (i item) View() string                        { return string(i) }

var _ = Describe("Selector", func() {
	Describe("when no items are given", func() {
		It("should render nothing", func() {
			m := wrapper{selector.New()}

			tm := teatest.NewTestModel(GinkgoTB(), m)
			tm.Send(tea.Quit())

			gtea.RequireGolden(tm)
			m = tm.FinalModel(GinkgoTB()).(wrapper)
			Expect(m.View()).To(BeEmpty())
		})
	})

	When("items are given", func() {
		It("should render the items", func() {
			m := wrapper{selector.New()}

			tm := teatest.NewTestModel(GinkgoTB(), m)
			tm.Send(selector.Items([]tea.Model{item("test")}))
			tm.Send(tea.Quit())

			gtea.RequireGolden(tm)
			m = tm.FinalModel(GinkgoTB()).(wrapper)
			Expect(m.View()).To(ContainSubstring("\u221f test"))
		})

		It("should render a cursor on the selected item", func() {
			m := wrapper{selector.New()}

			tm := teatest.NewTestModel(GinkgoTB(), m)
			tm.Send(selector.Items([]tea.Model{item("test1"), item("test2")}))
			tm.Send(tea.Quit())

			gtea.RequireGolden(tm)
			m = tm.FinalModel(GinkgoTB()).(wrapper)
			Expect(m.View()).To(ContainSubstring("\u221f test1 <"))
			Expect(m.View()).To(ContainSubstring("\u221f test2"))
		})

		It("should move the cursor when an arrow key is pressed", func() {
			m := wrapper{selector.New()}

			tm := teatest.NewTestModel(GinkgoTB(), m)
			tm.Send(selector.Items([]tea.Model{item("test1"), item("test2")}))
			tm.Send(tea.KeyMsg(tea.Key{Type: tea.KeyDown}))
			tm.Send(tea.Quit())

			gtea.RequireGolden(tm)
			m = tm.FinalModel(GinkgoTB()).(wrapper)
			Expect(m.View()).To(ContainSubstring("\u221f test1"))
			Expect(m.View()).To(ContainSubstring("\u221f test2 <"))
		})
	})
})
