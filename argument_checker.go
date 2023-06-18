package argument_checker

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// TODO: write description
const doc = "argument_checker is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "argument_checker",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if spec, ok := n.(*ast.TypeSpec); ok {
			if interfaceType, ok := spec.Type.(*ast.InterfaceType); ok {
				for _, method := range interfaceType.Methods.List {
					if funcType, ok := method.Type.(*ast.FuncType); ok {
						if len(funcType.Params.List) == 0 {
							pass.Reportf(method.Pos(), "context.Context is required as the first argument")
						} else {
							if selector, ok := funcType.Params.List[0].Type.(*ast.SelectorExpr); ok {
								if selector.Sel.Name != "Context" {
									pass.Reportf(method.Pos(), "context.Context is required as the first argument")
								}
							} else {
								pass.Reportf(method.Pos(), "context.Context is required as the first argument")
							}
						}
					}
				}
			}
		}
	})

	return nil, nil
}
