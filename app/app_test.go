package app_test

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/thecluster/app"
	"github.com/unmango/thecluster/project"
	"github.com/unmango/thecluster/testing/gtea"
)

var _ = Describe("App", Label("tea"), func() {
	It("should render project path", func() {
		m := app.New()
		m.Proj = &project.Project{Dir: work.Directory("/tests")}

		tm := teatest.NewTestModel(GinkgoTB(), m)
		tm.Send(tea.Quit())

		gtea.RequireGolden(tm)
		m = tm.FinalModel(GinkgoTB()).(app.Model)
		Expect(m.View()).To(ContainSubstring("/tests"))
	})

	It("should render errors", func() {
		m := app.New()
		m.Proj = &project.Project{}

		tm := teatest.NewTestModel(GinkgoTB(), m)
		tm.Send(fmt.Errorf("Test error"))
		tm.Send(tea.Quit())

		gtea.RequireGolden(tm)
	})
})
