package templates

import (
	"github.com/stretchr/testify/suite"
	"golang.stackrox.io/kube-linter/pkg/check"
	"golang.stackrox.io/kube-linter/pkg/diagnostic"
	"golang.stackrox.io/kube-linter/pkg/lintcontext"
)

// TemplateTestSuite is a basic testing suite for all templates
// test with some generic helper functions
type TemplateTestSuite struct {
	suite.Suite

	Template check.Template
}

// TestCase represents a single test case which can be verified under a LintContext
type TestCase struct {
	Param                    interface{}
	Diagnostics              map[string][]diagnostic.Diagnostic
	ExpectInstantiationError bool
}

// Init initializes the test suite with a template
func (s *TemplateTestSuite) Init(templateKey string) { _ = "STUB: not implemented"; return }

// Validate validates the given test cases against the LintContext passed in.
func (s *TemplateTestSuite) Validate(
	ctx lintcontext.LintContext,
	cases []TestCase,
) {
	_ = "STUB: not implemented"
	return
}

func (s *TemplateTestSuite) compareDiagnostics(expected, actual []diagnostic.Diagnostic) {
	_ = "STUB: not implemented"
	return
}
