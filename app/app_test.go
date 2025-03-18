package app_test

import (
	"fmt"
	"io"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/thecluster/app"
	"github.com/unmango/thecluster/project"
)

var _ = Describe("App", Label("tea"), func() {
	It("should render project path", func() {
		m := app.New()
		m.Proj = &project.Project{Dir: work.Directory("/test")}
		tm := teatest.NewTestModel(GinkgoTB(), m)
		tm.Send(tea.KeyCtrlC)

		result := tm.FinalOutput(GinkgoTB())

		out, err := io.ReadAll(result)
		Expect(err).NotTo(HaveOccurred())
		teatest.RequireEqualOutput(GinkgoTB(), out)
	})

	It("should render errors", func() {
		m := app.New()
		m.Update(fmt.Errorf("Test error"))
		tm := teatest.NewTestModel(GinkgoTB(), m)
		tm.Send(tea.KeyCtrlC)

		result := tm.FinalOutput(GinkgoTB())

		out, err := io.ReadAll(result)
		Expect(err).NotTo(HaveOccurred())
		teatest.RequireEqualOutput(GinkgoTB(), out)
	})
})
