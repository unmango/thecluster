package gtea

import (
	"io"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/charmbracelet/x/exp/teatest"
)

func RequireGolden(tm *teatest.TestModel) {
	ginkgo.GinkgoHelper()
	tb := ginkgo.GinkgoTB()

	out, err := io.ReadAll(tm.FinalOutput(tb))
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	teatest.RequireEqualOutput(tb, out)
}
