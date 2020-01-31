package visitors

import (
	"go/ast"
	"strings"
)

type FuncVisitor struct {
}

func (v *FuncVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch t := node.(type) {
	case *ast.FuncDecl:
		t.Name = ast.NewIdent(strings.Title(t.Name.Name))
	}

	return v
}
