package testing

import (
	"context"
	"os"
	"path/filepath"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
)

func PulumiWorkspace(ctx context.Context, appPath string) auto.Workspace {
	ginkgo.GinkgoHelper()

	gomega.Expect(os.Mkdir(appPath, os.ModePerm)).To(gomega.Succeed())
	work, err := auto.NewLocalWorkspace(ctx,
		auto.WorkDir(appPath),
		auto.Project(workspace.Project{
			Name:    tokens.PackageName(filepath.Dir(appPath)),
			Runtime: workspace.NewProjectRuntimeInfo("nodejs", nil),
		}),
	)
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	gomega.Expect(filepath.Join(appPath, "Pulumi.yaml")).To(gomega.BeARegularFile())

	return work
}
