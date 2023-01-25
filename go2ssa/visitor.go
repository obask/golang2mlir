package go2ssa

import (
	"awesomeProject/dialects/arith"
	"awesomeProject/dialects/fn"
	"awesomeProject/ir"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"
)

type GhostVisitor struct {
	fset         *token.FileSet
	name         string // Name of file.
	astFile      *ast.File
	Result       *ir.Operator
	pos          int
	currentBlock []ir.Operator
}

// Visit implements the ast.Visitor interface.
func (v *GhostVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return nil
	}
	typeName := fmt.Sprint(reflect.TypeOf(node))[5:]
	op := &ir.Operator{
		Name:        strings.ToLower(typeName[:1]) + typeName[1:],
		Dialect:     "go",
		Arguments:   nil,
		Blocks:      nil,
		ReturnNames: nil,
		Attributes:  map[string]ir.Attribute{},
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
		op.Attr("Imports", fmt.Sprintf("%v", n.Imports))
		newVisitor := GhostVisitor{}
		for _, decl := range n.Decls {
			ast.Walk(&newVisitor, decl)
		}
		op.Blocks = []ir.Block{{Items: newVisitor.currentBlock}}
		v.Result = op
		//op.Attr("Decls", fmt.Sprintf("%v", n.Decls))
		v.PushBack(op)
		return nil
	case *ast.FuncDecl:
		//		Doc  *CommentGroup // associated documentation; or nil
		//		Recv *FieldList    // receiver (methods); or nil (functions)
		//		Name *Ident        // function/method name
		//		Type *FuncType     // function signature: type and value parameters, results, and position of "fn" keyword
		//		Body *BlockStmt    // function body; or nil for external (non-Go) function
		op.T = ir.FuncFunc
		// TODO: add block arguments
		op.Attr("Recv", fmt.Sprintf("%v", n.Recv))
		op.Attributes["sym_name"] = ir.ReferenceAttr(fmt.Sprintf("%v", n.Name))
		op.Attr("Type", fmt.Sprintf("%v", n.Type))
		processRegion(op, n.Body)

		//newVisitor := GhostVisitor{}
		//ast.Walk(&newVisitor, n.Body)
		//op.Blocks = []ir.Block{{Items: newVisitor.currentBlock}}
		v.PushBack(op)
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
			op.Attr("name", f.Name)
		case *ast.SelectorExpr:
			switch lhs := f.X.(type) {
			case *ast.Ident:
				op = fn.MakeCallOp()
				op.Attributes["callee"] = ir.ReferenceAttr(lhs.Name + "." + f.Sel.Name)
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
			op.Attributes["value"] = ir.NumberAttr(ii)
			op.ReturnNames = []ir.SimpleName{ir.SimpleName("%c_" + n.Value)}
			op.ReturnTypes = []ir.SimpleType{"i32"}
			break
		default:
			op.Attr("kind", n.Kind.String())
			op.Attr("value", n.Value)
		}
		op.Attr("kind", n.Kind.String())
		op.Attr("value", n.Value)
	case *ast.GenDecl:
		//		Tok    token.Token   // IMPORT, CONST, TYPE, or VAR
		//		Specs  []Spec
		op.Attr("tok", n.Tok.String())
		op.Attr("specs", fmt.Sprintf("%v", n.Specs))
		v.PushBack(op)
		return nil
	case *ast.AssignStmt:
		//		Lhs    []Expr
		//		Tok    token.Token // assignment token, DEFINE
		//		Rhs    []Expr
		switch lhs := n.Lhs[0].(type) {
		case *ast.Ident:
			op.ReturnNames = append(op.ReturnNames, ir.SimpleName(lhs.Name))
			op.ReturnTypes = append(op.ReturnTypes, "interface{}")
			//op.Attr("var", lhs.Name)
		default:
			panic(lhs)
		}
		op.Attr("tok", n.Tok.String())
		v.processOperands(op, n.Rhs)
	case *ast.ForStmt:
		//		Init Stmt      // initialization statement; or nil
		//		Cond Expr      // condition; or nil
		//		Post Stmt      // post iteration statement; or nil
		//		Body *BlockStmt
		processRegion(op, n.Init)
		processRegion(op, n.Cond)
		processRegion(op, n.Post)
		processRegion(op, n.Body)
		v.PushBack(op)
		return nil
	case *ast.BinaryExpr:
		//		X     Expr        // left operand
		//		Op    token.Token // operator
		//		Y     Expr        // right operand
		op.Attr("op", n.Op.String())
		v.processOperands(op, []ast.Expr{n.X, n.Y})
	case *ast.Ident:
		//		Name    string    // identifier name
		//		Obj     *Object   // denoted object; or nil
		op.Attr("value", "%"+n.Name)
		op.Dialect = "mlir"
		op.Name = "load"
	case *ast.SelectorExpr:
		//		X   Expr   // expression
		//		Sel *Ident // field selector
		v.processOperands(op, []ast.Expr{n.X})
		op.Attr("sel", n.Sel.String())

		//println()
	case *ast.IfStmt:
		//		Init Stmt      // initialization statement; or nil
		//		Cond Expr      // condition
		//		Body *BlockStmt
		//		Else Stmt // else branch; or nil
		op.T = ir.ScfIf
		if n.Init != nil {
			panic("Init Stmt      // initialization statement; or nil")
		}
		v.processOperands(op, []ast.Expr{n.Cond})
		processRegion(op, n.Body)
		op.Blocks[0].Items = append(op.Blocks[0].Items, ir.Operator{T: ir.ScfYield})
		if n.Else != nil {
			processRegion(op, n.Else)
			op.Blocks[1].Items = append(op.Blocks[1].Items, ir.Operator{T: ir.ScfYield})
		}
		v.PushBack(op)
		return nil
	default:
		println("not found ->")
		fmt.Printf("case %T:\n", node)
		return nil
	}
	op.SetReturnName()
	v.PushBack(op)
	return nil
}

func processRegion(op *ir.Operator, stmt ast.Node) {
	newVisitor := GhostVisitor{}
	ast.Walk(&newVisitor, stmt)
	op.Blocks = append(op.Blocks, ir.Block{Items: newVisitor.currentBlock})
}

func (v *GhostVisitor) processOperands(op *ir.Operator, expressions []ast.Expr) {
	var newArguments []ir.ValueId
	for _, arg := range expressions {
		ast.Walk(v, arg)
		newArguments = append(newArguments, ir.SsaValue{Ref: &v.currentBlock[len(v.currentBlock)-1]})
	}
	//if len(newVisitor.currentBlock) != len(expressions) {
	//	panic("len(newVisitor.currentBlock) != len(expressions)")
	//}
	//tmp := ir.MapSlice(newArguments, func(o ir.Operator) ir.ValueId {
	//	//v.PushBack(&o)
	//	return ir.SsaValue{Ref: &o}
	//})
	op.Arguments = append(op.Arguments, newArguments...)
}

func (v *GhostVisitor) PushBack(operator *ir.Operator) {
	v.currentBlock = append(v.currentBlock, *operator)
}
