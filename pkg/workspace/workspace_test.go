package workspace_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/unmango/thecluster/pkg/project"
	"github.com/unmango/thecluster/pkg/workspace/pulumi"
)

var _ = Describe("Install", func() {
	var ws *pulumi.Workspace

	BeforeEach(func(ctx context.Context) {
		project, err := project.LocalGit(ctx)
		Expect(err).NotTo(HaveOccurred())
		path, err := project.Parse("cmd/testdata/test_install")
		Expect(err).NotTo(HaveOccurred())
		err = project.RemoveAll("cmd/testdata/test_install/node_modules")
		Expect(err).NotTo(HaveOccurred())
		ws, err = pulumi.Load(ctx, project, path)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should install dependencies", func(ctx context.Context) {
		err := ws.Install(ctx)

		Expect(err).NotTo(HaveOccurred())
	})
})
