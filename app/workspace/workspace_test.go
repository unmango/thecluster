package workspace_test

import (
	"context"
	"io"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/thecluster/app/workspace"
	"github.com/unmango/thecluster/project"
	"github.com/unmango/thecluster/testing"
	"github.com/unmango/thecluster/testing/gtea"
)

var _ = Describe("Workspace", func() {
	When("a Pulumi workspace exists", func() {
		var dir, appPath string

		BeforeEach(func(ctx context.Context) {
			dir = GinkgoT().TempDir()
			appPath = filepath.Join(dir, "app")
			testing.PulumiWorkspace(ctx, appPath)
		})

		It("should render loading", func(ctx context.Context) {
			m := workspace.New(ctx, project.Workspace(appPath))

			tm := teatest.NewTestModel(GinkgoTB(), m)
			tm.Send(tea.Quit())

			gtea.RequireGolden(tm)
			m = tm.FinalModel(GinkgoTB()).(workspace.Model)
			Expect(m.View()).To(ContainSubstring("Loading..."))
		})

		It("should render the workspace name", func(ctx context.Context) {
			m := workspace.New(ctx, project.Workspace(appPath))

			tm := teatest.NewTestModel(GinkgoTB(), m)
			Eventually(func() string {
				out, err := io.ReadAll(tm.Output())
				Expect(err).NotTo(HaveOccurred())
				return string(out)
			}).Should(ContainSubstring("app"))
			tm.Send(tea.Quit())

			gtea.RequireGolden(tm)
			m = tm.FinalModel(GinkgoTB()).(workspace.Model)
		})
	})
})
