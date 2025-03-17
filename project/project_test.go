package project_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/devctl/pkg/work"
	"github.com/unmango/thecluster/project"
)

var _ = Describe("Project", func() {
	It("should load the working directory", func(ctx context.Context) {
		expected, err := work.Load(ctx)
		Expect(err).NotTo(HaveOccurred())

		proj, err := project.Load(ctx)

		Expect(err).NotTo(HaveOccurred())
		Expect(proj.Dir).To(Equal(expected))
	})
})
