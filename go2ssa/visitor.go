package go2ssa

import (
	"awesomeProject/dialects/arith"
	"awesomeProject/hlir"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)
=
type GhostVisitor struct {
	fset           *token.FileSet
	name           string // Name of file.
	astFile        *ast.File
	Result         *hlir.Operator
	pos            int
	insertionPoint []hlir.Operator
}

// Visit implements the ast.Visitor interface.
func (v *GhostVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	typeName := fmt.Sprint(reflect.TypeOf(node))[5:]
	op := &hlir.Operator{
		Name:        strings.ToLower(typeName[:1]) + typeName[1:],
		Dialect:     "go",
		Arguments:   nil,
		Blocks:      nil,
		ReturnNames: nil,
		Attributes:  map[string]hlir.Attribute{},
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
		op.Attributes["Imports"] = hlir.StringAttr(fmt.Sprintf("%v", n.Imports))
		newVisitor := GhostVisitor{}
		for _, decl := range n.Decls {
			ast.Walk(&newVisitor, decl)
		}
		op.Blocks = []hlir.Block{{Items: newVisitor.insertionPoint}}
		v.Result = op
		//op.Attributes["Decls"] = hlir.StringAttr(fmt.Sprintf("%v", n.Decls))
		v.insertionPoint = append(v.insertionPoint, *op)
		return nil
	case *ast.FuncDecl:
		//		Doc  *CommentGroup // associated documentation; or nil
		//		Recv *FieldList    // receiver (methods); or nil (functions)
		//		Name *Ident        // function/method name
		//		Type *FuncType     // function signature: type and value parameters, results, and position of "func" keyword
		//		Body *BlockStmt    // function body; or nil for external (non-Go) function
		op.Attributes["Recv"] = hlir.StringAttr(fmt.Sprintf("%v", n.Recv))
		op.Attributes["Name"] = hlir.StringAttr(fmt.Sprintf("%v", n.Name))
		op.Attributes["Type"] = hlir.StringAttr(fmt.Sprintf("%v", n.Type))
		newVisitor := GhostVisitor{}
		ast.Walk(&newVisitor, n.Body)
		op.Blocks = []hlir.Block{{Items: newVisitor.insertionPoint}}
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
			op.Attributes["name"] = hlir.StringAttr(f.Name)
		case *ast.SelectorExpr:
			switch lhs := f.X.(type) {
			case *ast.Ident:
				op.Attributes["name"] = hlir.StringAttr(lhs.Name + "." + f.Sel.Name)
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
		switch n.Kind {
		case token.INT:
			op = arith.MakeConstant()
			ii, _ := strconv.Atoi(n.Value)
			op.Attributes["value"] = hlir.NumberAttr(ii)
			op.ReturnNames = []hlir.ValueId{hlir.ValueId("%c_" + n.Value)}
			op.ReturnTypes = []hlir.SimpleType{"i32"}
			break
		default:
			op.Attributes["kind"] = hlir.StringAttr(n.Kind.String())
			op.Attributes["value"] = hlir.StringAttr(n.Value)
		}
		op.Attributes["kind"] = hlir.StringAttr(n.Kind.String())
		op.Attributes["value"] = hlir.StringAttr(n.Value)
	case *ast.GenDecl:
		//		Tok    token.Token   // IMPORT, CONST, TYPE, or VAR
		//		Specs  []Spec
		op.Attributes["tok"] = hlir.StringAttr(n.Tok.String())
		op.Attributes["specs"] = hlir.StringAttr(fmt.Sprintf("%v", n.Specs))
		v.insertionPoint = append(v.insertionPoint, *op)
		return nil
	case *ast.AssignStmt:
		//		Lhs    []Expr
		//		Tok    token.Token // assignment token, DEFINE
		//		Rhs    []Expr
		switch lhs := n.Lhs[0].(type) {
		case *ast.Ident:
			op.Attributes["var"] = hlir.StringAttr(lhs.Name)
		default:
			panic(lhs)
		}
		op.Attributes["tok"] = hlir.StringAttr(n.Tok.String())
		v.processOperands(op, n.Rhs)
	case *ast.ForStmt:
		//		Init Stmt      // initialization statement; or nil
		//		Cond Expr      // condition; or nil
		//		Post Stmt      // post iteration statement; or nil
		//		Body *BlockStmt
		processRegion(n.Init, op)
		processRegion(n.Cond, op)
		processRegion(n.Post, op)
		processRegion(n.Body, op)
		v.insertionPoint = append(v.insertionPoint, *op)
		return nil
	case *ast.BinaryExpr:
		//		X     Expr        // left operand
		//		Op    token.Token // operator
		//		Y     Expr        // right operand
		op.Attributes["op"] = hlir.StringAttr(n.Op.String())
		v.processOperands(op, []ast.Expr{n.X, n.Y})
	case *ast.Ident:
		//		Name    string    // identifier name
		//		Obj     *Object   // denoted object; or nil
		op.Attributes["name"] = hlir.StringAttr(n.Name)
	case *ast.SelectorExpr:
		//		X   Expr   // expression
		//		Sel *Ident // field selector
		v.processOperands(op, []ast.Expr{n.X})
		op.Attributes["sel"] = hlir.StringAttr(n.Sel.String())

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

func processRegion(stmt ast.Node, op *hlir.Operator) {
	newVisitor := GhostVisitor{}
	ast.Walk(&newVisitor, stmt)
	op.Blocks = append(op.Blocks, hlir.Block{Items: newVisitor.insertionPoint})
}

func (v *GhostVisitor) processOperands(op *hlir.Operator, expressions []ast.Expr) {
	newVisitor := &GhostVisitor{}
	for _, arg := range expressions {
		//fmt.Printf("%+v\n", arg)
		ast.Walk(newVisitor, arg)
	}
	v.insertionPoint = append(v.insertionPoint, newVisitor.insertionPoint...)
	for _, operand := range newVisitor.insertionPoint {
		op.Arguments = append(op.Arguments, operand.ReturnNames...)
	}
}
