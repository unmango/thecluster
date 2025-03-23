package project_test

import (
	"context"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/thecluster/project"
	"github.com/unmango/thecluster/testing"
)

var _ = Describe("Workspace", func() {
	When("the workspace is a Pulumi project", func() {
		var dir, appPath string

		BeforeEach(func(ctx context.Context) {
			dir = GinkgoT().TempDir()
			appPath = filepath.Join(dir, "app")
			testing.PulumiWorkspace(ctx, appPath)
		})

		It("should load the Pulumi workspace", func(ctx context.Context) {
			w := project.Workspace(appPath)

			work, err := w.Load(ctx)

			Expect(err).NotTo(HaveOccurred())
			Expect(work).NotTo(BeNil())
		})
	})
})
