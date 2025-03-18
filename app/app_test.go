package app_test

import (
	"io"

	"github.com/charmbracelet/x/exp/teatest/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unmango/thecluster/app"
)

var _ = Describe("App", Label("tea"), func() {
	It("should render", func() {
		m := &app.Model{}
		tm := teatest.NewTestModel(GinkgoTB(), m)

		result := tm.FinalOutput(GinkgoTB())

		out, err := io.ReadAll(result)
		Expect(err).NotTo(HaveOccurred())
		teatest.RequireEqualOutput(GinkgoTB(), out)
	})
})
