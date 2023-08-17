package argument_checker_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/toshiki-otaka/argument_checker"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, argument_checker.Analyzer, "query", "command")
}
