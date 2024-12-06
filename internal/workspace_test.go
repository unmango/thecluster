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
	Describe("LoadWorkspace", func() {
		var fsys afero.Fs

		BeforeEach(func() {
			fsys = afero.NewMemMapFs()
		})

		It("should load", func() {
			err := fsys.Mkdir("test", os.ModeDir)
			Expect(err).NotTo(HaveOccurred())

			ws, err := internal.LoadWorkspace(fsys, "test")

			Expect(err).NotTo(HaveOccurred())
			Expect(ws.Name()).To(Equal("test"))
		})
	})

	Describe("Parse", func() {
		var ws pkg.Workspace

		BeforeEach(func() {
			fsys := afero.NewMemMapFs()
			err := fsys.Mkdir("test", os.ModeDir)
			Expect(err).NotTo(HaveOccurred())
			ws, err = internal.LoadWorkspace(fsys, "test")
			Expect(err).NotTo(HaveOccurred())
		})

		It("should return a relative path", func() {
			fn := func(p string) bool {
				r, err := ws.Parse(p)

				return err == nil && !filepath.IsAbs(r)
			}

			Expect(quick.Check(fn, nil)).To(Succeed())
		})
	})
})
