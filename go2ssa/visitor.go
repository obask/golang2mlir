package go2ssa

import (
	"awesomeProject/mlir"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strings"
)

type GhostVisitor struct {
	fset           *token.FileSet
	name           string // Name of file.
	astFile        *ast.File
	Result         *mlir.Operator
	pos            int
	insertionPoint []mlir.Operator
}

// Visit implements the ast.Visitor interface.
func (v *GhostVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	typeName := fmt.Sprint(reflect.TypeOf(node))[5:]
	op := &mlir.Operator{
		Name:       strings.ToLower(typeName[:1]) + typeName[1:],
		Dialect:    "go",
		Operands:   nil,
		Regions:    nil,
		ReturnName: "",
		Attributes: map[string]mlir.Attribute{},
	}
	switch n := node.(type) {
	case *ast.File:
		//op.Name = "file"
		//	Doc        *CommentGroup   // associated documentation; or nil
		//	Name       *Ident          // package name
		//	Decls      []Decl          // top-level declarations; or nil
		//	Scope      *Scope          // package scope (this file only)
		//	Imports    []*ImportSpec   // imports in this file
		//	Unresolved []*Ident        // unresolved identifiers in this file
		//	Comments   []*CommentGroup // list of all comments in the source file
		op.Attributes["Imports"] = mlir.StringAttr(fmt.Sprintf("%v", n.Imports))
		newVisitor := GhostVisitor{}
		for _, decl := range n.Decls {
			ast.Walk(&newVisitor, decl)
		}
		op.Regions = []mlir.Region{[]mlir.BasicBlock{{Items: newVisitor.insertionPoint}}}
		v.Result = op
		//op.Attributes["Decls"] = mlir.StringAttr(fmt.Sprintf("%v", n.Decls))
		v.insertionPoint = append(v.insertionPoint, *op)
		return nil
	case *ast.FuncDecl:
		//		Doc  *CommentGroup // associated documentation; or nil
		//		Recv *FieldList    // receiver (methods); or nil (functions)
		//		Name *Ident        // function/method name
		//		Type *FuncType     // function signature: type and value parameters, results, and position of "func" keyword
		//		Body *BlockStmt    // function body; or nil for external (non-Go) function
		op.Attributes["Recv"] = mlir.StringAttr(fmt.Sprintf("%v", n.Recv))
		op.Attributes["Name"] = mlir.StringAttr(fmt.Sprintf("%v", n.Name))
		op.Attributes["Type"] = mlir.StringAttr(fmt.Sprintf("%v", n.Type))
		newVisitor := GhostVisitor{}
		ast.Walk(&newVisitor, n.Body)
		op.Regions = []mlir.Region{[]mlir.BasicBlock{{Items: newVisitor.insertionPoint}}}
		v.insertionPoint = append(v.insertionPoint, *op)
		return nil
	case *ast.BlockStmt:
		//		List   []Stmt
		return v
	case *ast.ExprStmt:
		//		X Expr // expression
		return v
	case *ast.CallExpr:
		//		Fun      Expr      // function expression
		//		Args     []Expr    // function arguments; or nil
		switch f := n.Fun.(type) {
		case *ast.Ident:
			op.Attributes["name"] = mlir.StringAttr(f.Name)
		case *ast.SelectorExpr:
			switch lhs := f.X.(type) {
			case *ast.Ident:
				op.Attributes["name"] = mlir.StringAttr(lhs.Name + "." + f.Sel.Name)
			default:
				panic(lhs)
			}
		default:
			panic(f)
		}
		v.processOperands(op, n.Args)
		op.SetReturnName()
	case *ast.BasicLit:
		//Kind     token.Token // token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING
		//Value    string      // literal string; e.g. 42, 0x7f, 3.14, 1e-9, 2.4i, 'a', '\x7f', "foo" or `\m\n\o`
		op.Attributes["kind"] = mlir.StringAttr(n.Kind.String())
		op.Attributes["value"] = mlir.StringAttr(n.Value)
	case *ast.GenDecl:
		//		Tok    token.Token   // IMPORT, CONST, TYPE, or VAR
		//		Specs  []Spec
		op.Attributes["tok"] = mlir.StringAttr(n.Tok.String())
		op.Attributes["specs"] = mlir.StringAttr(fmt.Sprintf("%v", n.Specs))
		v.insertionPoint = append(v.insertionPoint, *op)
		return nil
	case *ast.AssignStmt:
		//		Lhs    []Expr
		//		Tok    token.Token // assignment token, DEFINE
		//		Rhs    []Expr
		switch lhs := n.Lhs[0].(type) {
		case *ast.Ident:
			op.Attributes["var"] = mlir.StringAttr(lhs.Name)
		default:
			panic(lhs)
		}
		op.Attributes["tok"] = mlir.StringAttr(n.Tok.String())
		v.processOperands(op, n.Rhs)
	case *ast.ForStmt:
		//		Init Stmt      // initialization statement; or nil
		//		Cond Expr      // condition; or nil
		//		Post Stmt      // post iteration statement; or nil
		//		Body *BlockStmt
		newVisitor := GhostVisitor{}
		ast.Walk(&newVisitor, n.Init)
		op.Regions = append(op.Regions, []mlir.BasicBlock{{Items: newVisitor.insertionPoint}})
		newVisitor = GhostVisitor{}
		ast.Walk(&newVisitor, n.Cond)
		op.Regions = append(op.Regions, []mlir.BasicBlock{{Items: newVisitor.insertionPoint}})
		newVisitor = GhostVisitor{}
		ast.Walk(&newVisitor, n.Post)
		op.Regions = append(op.Regions, []mlir.BasicBlock{{Items: newVisitor.insertionPoint}})
		newVisitor = GhostVisitor{}
		ast.Walk(&newVisitor, n.Body)
		op.Regions = append(op.Regions, []mlir.BasicBlock{{Items: newVisitor.insertionPoint}})
		v.insertionPoint = append(v.insertionPoint, *op)
		return nil
	case *ast.BinaryExpr:
		//		X     Expr        // left operand
		//		Op    token.Token // operator
		//		Y     Expr        // right operand
		op.Attributes["op"] = mlir.StringAttr(n.Op.String())
		v.processOperands(op, []ast.Expr{n.X, n.Y})
	case *ast.Ident:
		//		Name    string    // identifier name
		//		Obj     *Object   // denoted object; or nil
		op.Attributes["name"] = mlir.StringAttr(n.Name)
	case *ast.SelectorExpr:
		//		X   Expr   // expression
		//		Sel *Ident // field selector
		v.processOperands(op, []ast.Expr{n.X})
		op.Attributes["sel"] = mlir.StringAttr(n.Sel.String())

		println()
	default:
		println("not found ->")
		fmt.Printf("case %T:\n", node)
		return nil
	}
	op.SetReturnName()
	v.insertionPoint = append(v.insertionPoint, *op)
	return nil
}

func (v *GhostVisitor) processOperands(op *mlir.Operator, expressions []ast.Expr) {
	newVisitor := &GhostVisitor{}
	for _, arg := range expressions {
		//fmt.Printf("%+v\n", arg)
		ast.Walk(newVisitor, arg)
	}
	v.insertionPoint = append(v.insertionPoint, newVisitor.insertionPoint...)
	for _, operand := range newVisitor.insertionPoint {
		op.Operands = append(op.Operands, operand.ReturnName)
	}
}
