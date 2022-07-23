package go2ssa

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
	Result  *mlir.Operator
	pos     int
}

// Visit implements the ast.Visitor interface.
func (v *GhostVisitor) Visit(node ast.Node) ast.Visitor {
	op := &mlir.Operator{
		Name:       "",
		Dialect:    "",
		Operands:   nil,
		Regions:    []mlir.Region{},
		ReturnName: "",
		Attributes: map[string]mlir.Attribute{},
	}
	switch n := node.(type) {
	case *ast.File:
		//	Doc        *CommentGroup   // associated documentation; or nil
		//	Package    token.Pos       // position of "package" keyword
		//	Name       *Ident          // package name
		//	Decls      []Decl          // top-level declarations; or nil
		//	Scope      *Scope          // package scope (this file only)
		//	Imports    []*ImportSpec   // imports in this file
		//	Unresolved []*Ident        // unresolved identifiers in this file
		//	Comments   []*CommentGroup // list of all comments in the source file
		op.Attributes["Imports"] = mlir.StringAttr(fmt.Sprintf("%v", n.Imports))
		insertionPoint := mlir.BasicBlock{
			Label: mlir.BlockLabel{},
			Items: nil,
		}
		// TODO invoke walk for each Decl
		for i, decl := range n.Decls {
			insertionPoint.Items = append(insertionPoint.Items)
		}

		op.Attributes["Decls"] = mlir.StringAttr(fmt.Sprintf("%v", n.Decls))

	case *ast.FuncDecl:
		//		Doc  *CommentGroup // associated documentation; or nil
		//		Recv *FieldList    // receiver (methods); or nil (functions)
		//		Name *Ident        // function/method name
		//		Type *FuncType     // function signature: type and value parameters, results, and position of "func" keyword
		//		Body *BlockStmt    // function body; or nil for external (non-Go) function
		op.Attributes["Recv"] = mlir.StringAttr(fmt.Sprintf("%v", n.Recv))
		op.Attributes["Name"] = mlir.StringAttr(fmt.Sprintf("%v", n.Name))
		op.Attributes["Type"] = mlir.StringAttr(fmt.Sprintf("%v", n.Type))
		op.Attributes["Body"] = mlir.StringAttr(fmt.Sprintf("%v", n.Body))
	default:
		fmt.Printf("%T\n", node)
	}

	v.Result = op
	return nil
}
