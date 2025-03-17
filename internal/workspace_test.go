package internal_test

import (
	"os"
	"path/filepath"
	"testing/quick"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/unmango/thecluster/internal"
	"github.com/unmango/thecluster/pkg"
)

var _ = Describe("Workspace", func() {
	Describe("Parse", func() {
		var ws pkg.Workspace

		BeforeEach(func() {
			fsys := afero.NewMemMapFs()
			err := fsys.Mkdir("/test", os.ModeDir)
			Expect(err).NotTo(HaveOccurred())
			ws = internal.Workspace{fsys, "/test"}
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return an absolute path", func() {
			fn := func(p string) bool {
				r, err := ws.Parse(filepath.Join("/test", p))

				return err == nil && filepath.IsAbs(r)
			}

			Expect(quick.Check(fn, nil)).To(Succeed())
		})

		It("should join to root", func() {
			r, err := ws.Parse("subdir")

			Expect(err).NotTo(HaveOccurred())
			Expect(r).To(Equal("/test/subdir"))
		})

		It("should return a rooted path", func() {
			r, err := ws.Parse("/test/subdir")

			Expect(err).NotTo(HaveOccurred())
			Expect(r).To(Equal("/test/subdir"))
		})

		It("should join a non-matching path", func() {
			r, err := ws.Parse("/subdir")

			Expect(err).NotTo(HaveOccurred())
			Expect(r).To(Equal("/test/subdir"))
		})
	})
})
