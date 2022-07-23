package goast2ssa

import (
	"awesomeProject/mlir"
	"fmt"
	"go/ast"
	"go/token"
)

type GhostVisitor struct {
	fset    *token.FileSet
	name    string // Name of file.
	astFile *ast.File
	funcs   []*mlir.Operator
	pos     int
}

// Visit implements the ast.Visitor interface.
func (v *GhostVisitor) Visit(node ast.Node) ast.Visitor {
	fmt.Printf("%d: %+v", v.pos, node)
	switch n := node.(type) {
	case *ast.FuncDecl:
		if n.Body == nil {
			// Do not count declarations of assembly functions.
			break
		}
		fe := &mlir.Operator{
			Name:       "",
			Dialect:    "",
			Operands:   nil,
			Regions:    nil,
			ReturnName: "",
			Attributes: nil,
		}
		v.funcs = append(v.funcs, fe)
	}
	return &GhostVisitor{pos: v.pos + 1}
}
