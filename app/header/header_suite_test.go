package header_test

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHeader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Header Suite")
}

var _ = BeforeSuite(func() {
	lipgloss.SetColorProfile(termenv.Ascii)
})
