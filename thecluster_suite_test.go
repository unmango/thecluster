package main_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var cliPath string

func TestThecluster(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Thecluster Suite")
}

var _ = BeforeSuite(func() {
	wd, err := os.Getwd()
	Expect(err).NotTo(HaveOccurred())
	cliPath, err = gexec.Build(wd)
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
