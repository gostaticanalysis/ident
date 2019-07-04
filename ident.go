package ident

import (
	"go/ast"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer maps types.Object to *ast.Ident.
var Analyzer = &analysis.Analyzer{
	Name: "ident",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	ResultType: reflect.TypeOf(Map{}),
}

const doc = "ident maps types.Object to *ast.Ident"

// Map maps types.Object to *ast.Ident.
type Map map[types.Object][]*ast.Ident

func run(pass *analysis.Pass) (interface{}, error) {
	m := Map{}
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.Ident:
			if n.Name != "_" {
				if o := pass.TypesInfo.ObjectOf(n); o != nil {
					m[o] = append(m[o], n)
				}
			}
		}
	})

	return m, nil
}
