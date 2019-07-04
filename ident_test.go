package ident_test

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/gostaticanalysis/ident"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := &analysis.Analyzer{
		Requires: []*analysis.Analyzer{
			ident.Analyzer,
		},
		Run: func(pass *analysis.Pass) (interface{}, error) {
			m := pass.ResultOf[ident.Analyzer].(ident.Map)
			for o, idents := range m {
				sort.Slice(idents, func(i, j int) bool {
					return idents[i].Pos() < idents[j].Pos()
				})
				lines := make([]string, len(idents))
				for i := range idents {
					pos := pass.Fset.Position(idents[i].Pos())
					lines[i] = fmt.Sprintf("%s:%d", filepath.Base(pos.Filename), pos.Line)
				}
				pass.Reportf(o.Pos(), "%s", strings.Join(lines, " "))
			}
			return nil, nil
		},
	}
	analysistest.Run(t, testdata, analyzer, "a")
}
