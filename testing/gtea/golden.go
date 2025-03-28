package gtea

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/aymanbagabas/go-udiff"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/charmbracelet/x/exp/teatest/v2"
)

func RequireGolden(tm *teatest.TestModel) {
	ginkgo.GinkgoHelper()
	tb := ginkgo.GinkgoTB()

	out, err := io.ReadAll(tm.FinalOutput(tb))
	gomega.Expect(err).NotTo(gomega.HaveOccurred())
	teatest.RequireEqualOutput(tb, out)
}

type golden struct {
	tb testing.TB

	goldenStr string
	outStr    string
	diff      string
}

func BeGolden(tb testing.TB) *golden {
	return &golden{tb, "", "", ""}
}

// Match implements types.GomegaMatcher.
func (g *golden) Match(actual interface{}) (success bool, err error) {
	out, ok := actual.([]byte)
	if !ok {
		return false, fmt.Errorf("expected a []byte, got %v", reflect.TypeOf(actual))
	}

	golden := filepath.Join("testdata", g.tb.Name()+".golden")
	goldenBts, err := os.ReadFile(golden)
	if err != nil {
		return false, err
	}

	g.goldenStr = normalizeWindowsLineBreaks(string(goldenBts))
	g.goldenStr = escapeSeqs(g.goldenStr)
	g.outStr = escapeSeqs(string(out))

	g.diff = udiff.Unified("golden", "run", g.goldenStr, g.outStr)
	return g.diff == "", nil
}

// FailureMessage implements types.GomegaMatcher.
func (g *golden) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf(
		"output does not match, expected:\n\n%s\n\ngot:\n\n%s\n\ndiff:\n\n%s",
		g.goldenStr,
		g.outStr,
		g.diff,
	)
}

// NegatedFailureMessage implements types.GomegaMatcher.
func (g *golden) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("output matched, expected a diff: %s", g.outStr)
}

// https://github.com/charmbracelet/x/blob/2fdc97757edf95bb82d27567342f99167b3e71ab/exp/golden/golden.go#L64
func escapeSeqs(in string) string {
	s := strings.Split(in, "\n")
	for i, l := range s {
		q := strconv.Quote(l)
		q = strings.TrimPrefix(q, `"`)
		q = strings.TrimSuffix(q, `"`)
		s[i] = q
	}
	return strings.Join(s, "\n")
}

// https://github.com/charmbracelet/x/blob/2fdc97757edf95bb82d27567342f99167b3e71ab/exp/golden/golden.go#L77
func normalizeWindowsLineBreaks(str string) string {
	if runtime.GOOS == "windows" {
		return strings.ReplaceAll(str, "\r\n", "\n")
	}
	return str
}
