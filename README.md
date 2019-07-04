# ident [![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godoc] [![Travis](https://img.shields.io/travis/gostaticanalysis/ident.svg?style=flat-square)][travis] [![Go Report Card](https://goreportcard.com/badge/github.com/gostaticanalysis/ident)](https://goreportcard.com/report/github.com/gostaticanalysis/ident) [![codecov](https://codecov.io/gh/gostaticanalysis/ident/branch/master/graph/badge.svg)](https://codecov.io/gh/gostaticanalysis/ident)

`ident` maps `types.Object` to `*ast.Ident`.

## Usage

```go
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
```

<!-- links -->
[godoc]: http://godoc.org/github.com/gostaticanalysis/ident
[travis]: https://travis-ci.org/gostaticanalysis/ident
