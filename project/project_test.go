package project_test

import (
	"context"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"

	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/thecluster/project"
)

var _ = Describe("Project", func() {
	Describe("Load", func() {
		It("should load the working directory", func(ctx context.Context) {
			expected, err := work.Load(ctx)
			Expect(err).NotTo(HaveOccurred())

			proj, err := project.Load(ctx)

			Expect(err).NotTo(HaveOccurred())
			Expect(proj.Dir).To(Equal(expected))
		})
	})

	Describe("LoadFrom", func() {
		It("should load the given path", func() {
			path := GinkgoT().TempDir()

			proj, err := project.LoadFrom(path)

			Expect(err).NotTo(HaveOccurred())
			Expect(proj.Dir.Path()).To(Equal(path))
		})

		It("should error when path does not exist", func() {
			path := filepath.Join(GinkgoT().TempDir(), "blah")

			_, err := project.LoadFrom(path)

			Expect(err).To(MatchError(
				ContainSubstring("no such file or directory"),
			))
		})
	})

	Describe("Workspaces", func() {
		It("should return an empty seq", func(ctx context.Context) {
			proj, err := project.LoadFrom(GinkgoT().TempDir())

			Expect(err).NotTo(HaveOccurred())
			ws, err := proj.Workspaces()
			Expect(err).NotTo(HaveOccurred())
			Expect(ws).To(BeEmpty())
		})

		When("a Pulumi workspace exists", func() {
			var dir, appPath string

			BeforeEach(func(ctx context.Context) {
				dir = GinkgoT().TempDir()
				appPath = filepath.Join(dir, "app")
				Expect(os.Mkdir(appPath, os.ModePerm)).To(Succeed())

				_, err := auto.NewLocalWorkspace(ctx,
					auto.WorkDir(appPath),
					auto.Project(workspace.Project{
						Name:    "TEST",
						Runtime: workspace.NewProjectRuntimeInfo("nodejs", nil),
					}),
				)
				Expect(err).NotTo(HaveOccurred())
				Expect(filepath.Join(appPath, "Pulumi.yaml")).To(BeARegularFile())
			})

			It("should list the workspace", func(ctx context.Context) {
				proj, err := project.LoadFrom(dir)

				Expect(err).NotTo(HaveOccurred())
				ws, err := proj.Workspaces()
				Expect(err).NotTo(HaveOccurred())
				Expect(ws).To(ConsistOf(appPath))
			})
		})
	})
})
