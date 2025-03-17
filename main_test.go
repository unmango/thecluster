package main_test

import (
	"context"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	It("should load a project at the current filepath", func(ctx context.Context) {
		wd, err := os.Getwd()
		Expect(err).NotTo(HaveOccurred())
		cmd := exec.CommandContext(ctx, cliPath)

		ses, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())
		Eventually(ses.Out).Should(gbytes.Say("Project: "))
		Eventually(ses.Out).Should(gbytes.Say(wd))
		Eventually(ses).Should(gexec.Exit(0))
	})
})
