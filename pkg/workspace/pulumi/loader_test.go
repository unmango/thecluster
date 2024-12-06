package pulumi_test

import (
	"context"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/unmango/thecluster/pkg/testing"
	"github.com/unmango/thecluster/pkg/workspace/pulumi"
)

var _ = Describe("Loader", func() {
	DescribeTable("should load a pulumi workspace",
		Entry(nil, "Pulumi.yaml"),
		Entry(nil, "Pulumi.yml"),
		func(ctx context.Context, f string) {
			fs := afero.NewMemMapFs()
			err := fs.Mkdir("test", os.ModeDir)
			Expect(err).NotTo(HaveOccurred())
			_, err = fs.Create(filepath.Join("test", f))
			Expect(err).NotTo(HaveOccurred())
			pctx := testing.DefaultContext(ctx, fs)

			w, err := pulumi.Loader.Load(pctx, "test")

			Expect(err).NotTo(HaveOccurred())
			Expect(w.Name()).To(Equal("test"))
		},
	)

	It("should NOT load an unsupported workspace", func(ctx context.Context) {
		fs := afero.NewMemMapFs()
		err := fs.Mkdir("test", os.ModeDir)
		Expect(err).NotTo(HaveOccurred())
		_, err = fs.Create("test/Blah.yaml")
		Expect(err).NotTo(HaveOccurred())
		pctx := testing.DefaultContext(ctx, fs)

		_, err = pulumi.Loader.Load(pctx, "test")

		Expect(err).To(MatchError("not a pulumi workspace: test"))
	})
})
