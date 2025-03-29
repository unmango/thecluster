package app_test

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/x/exp/teatest/v2"
	. "github.com/onsi/ginkgo/v2"

	"github.com/unmango/thecluster/app"
	"github.com/unmango/thecluster/testing/gtea"
)

var _ = Describe("App", func() {
	It("should render", func() {
		m := app.New()

		tm := teatest.NewTestModel(GinkgoTB(), m)
		tm.Send(tea.Quit())

		gtea.RequireGolden(tm)
	})
})
