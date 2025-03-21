package header_test

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/thecluster/app/header"
	"github.com/unmango/thecluster/testing/gtea"
)

var _ = Describe("Header", func() {
	It("should render", func() {
		m := header.New()
		m.Title = "TEST"

		tm := teatest.NewTestModel(GinkgoTB(), m)
		tm.Send(tea.Quit())

		gtea.RequireGolden(tm)
		m = tm.FinalModel(GinkgoTB()).(header.Model)
		Expect(m.Title).To(Equal("TEST"))
	})
})
