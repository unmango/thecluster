package pulumi_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPulumi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pulumi Suite")
}
