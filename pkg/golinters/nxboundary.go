package golinters

import (
	"fmt"
	"strings"

	"github.com/vanstinator/nxboundary"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/golinters/goanalysis"
	"github.com/golangci/golangci-lint/pkg/lint/linter"
)

func NewNxBoundary(settings *config.NxBoundarySettings) *goanalysis.Linter {
	analyzer := nxboundary.NewAnalyzer()

	return goanalysis.NewLinter(
		analyzer.Name,
		analyzer.Doc,
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithContextSetter(func(lintCtx *linter.Context) {
		if settings == nil {
			return
		}
		if len(settings.AllowedTags) == 0 {
			lintCtx.Log.Infof("nxboundary settings found, but no allowedTags listed. List aliases under alias: key.") //nolint:misspell
		}

		for key, value := range settings.AllowedTags {
			err := analyzer.Flags.Set("allowedTags", fmt.Sprintf("%s|%s", key, strings.Join(value, ",")))
			if err != nil {
				lintCtx.Log.Errorf("failed to parse configuration: %v", err)
			}
		}
	}).WithLoadMode(goanalysis.LoadModeSyntax)
}
