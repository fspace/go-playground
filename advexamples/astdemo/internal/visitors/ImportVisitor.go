package visitors

import (
	"go/ast"
	"go/token"
)

type ImportVisitor struct{}

func (i *ImportVisitor) Visit(node ast.Node) (w ast.Visitor) {
	switch t := node.(type) {
	case *ast.GenDecl:
		if t.Tok == token.IMPORT {
			newSpecs := make([]ast.Spec, len(t.Specs)+1)
			for i, spec := range t.Specs {
				newSpecs[i] = spec
			}
			newPackage := &ast.BasicLit{token.NoPos, token.STRING, "fandango"}
			newSpecs[len(t.Specs)] = &ast.ImportSpec{
				Doc:     nil,
				Name:    nil,
				Path:    newPackage,
				Comment: nil,
			}
			t.Specs = newSpecs
		}
		return nil
	}

	return i
}
