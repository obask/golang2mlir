package go2ssa

import (
	"awesomeProject/hlir"
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strings"
)

type GhostConvertor struct {
	fset           *token.FileSet
	name           string // Name of file.
	astFile        *ast.File
	Result         *hlir.Operator
	pos            int
	insertionPoint [][]hlir.Operator
}

func makeGoOperator(node ast.Node) hlir.Operator {
	typeName := fmt.Sprint(reflect.TypeOf(node))[5:]
	return hlir.Operator{
		Name:    strings.ToLower(typeName[:1]) + typeName[1:],
		Dialect: "go",
	}
}

func (c *GhostConvertor) processRegion(op *hlir.Operator, items []hlir.Operator) {
	op.Blocks = append(op.Blocks, hlir.Block{Items: items})
}

func (c *GhostConvertor) processOperands(m *hlir.Operator, args []ast.Expr) {
	c.insertionPoint = append(c.insertionPoint, nil)

	c.insertionPoint = c.insertionPoint[:len(c.insertionPoint)-1]
}

func (c *GhostConvertor) visitFile(file *ast.File) hlir.Operator {
	op := makeGoOperator(file)
	op.Attributes["Imports"] = hlir.StringAttr(fmt.Sprintf("%v", file.Imports))
	c.processRegion(&op, c.visitDeclarations(file.Decls))
	newVisitor := GhostVisitor{}
	for _, decl := range file.Decls {
		ast.Walk(&newVisitor, decl)
	}

	return op
}

func (c *GhostConvertor) visitBadDecl(decl *ast.BadDecl) hlir.Operator {
	op := makeGoOperator(decl)
	return op
}

func (c *GhostConvertor) visitGenDecl(decl *ast.GenDecl) hlir.Operator {
	op := makeGoOperator(decl)
	return op
}

func (c *GhostConvertor) visitFuncDecl(decl *ast.FuncDecl) hlir.Operator {
	op := makeGoOperator(decl)
	op.Attributes["Recv"] = hlir.StringAttr(fmt.Sprintf("%v", decl.Recv))
	op.Attributes["Name"] = hlir.StringAttr(fmt.Sprintf("%v", decl.Name))
	op.Attributes["Type"] = hlir.StringAttr(fmt.Sprintf("%v", decl.Type))
	c.processRegion(&op, c.visitStatements(decl.Body.List))
	return op
}

func (c *GhostConvertor) visitArrayType(expr *ast.ArrayType) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitBadExpr(expr *ast.BadExpr) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitBasicLit(expr *ast.BasicLit) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitBinaryExpr(expr *ast.BinaryExpr) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitCallExpr(expr *ast.CallExpr) hlir.Operator {
	op := makeGoOperator(expr)
	switch f := expr.Fun.(type) {
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
	c.processOperands(&op, expr.Args)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitChanType(expr *ast.ChanType) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitCompositeLit(expr *ast.CompositeLit) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitEllipsis(expr *ast.Ellipsis) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitFuncLit(expr *ast.FuncLit) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitFuncType(expr *ast.FuncType) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitIdent(expr *ast.Ident) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitIndexExpr(expr *ast.IndexExpr) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitIndexListExpr(expr *ast.IndexListExpr) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitInterfaceType(expr *ast.InterfaceType) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitKeyValueExpr(expr *ast.KeyValueExpr) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitMapType(expr *ast.MapType) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitParenExpr(expr *ast.ParenExpr) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitSelectorExpr(expr *ast.SelectorExpr) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitSliceExpr(expr *ast.SliceExpr) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitStarExpr(expr *ast.StarExpr) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitStructType(expr *ast.StructType) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitTypeAssertExpr(expr *ast.TypeAssertExpr) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitUnaryExpr(expr *ast.UnaryExpr) hlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitAssignStmt(stmt *ast.AssignStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitBadStmt(stmt *ast.BadStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitBlockStmt(stmt *ast.BlockStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitBranchStmt(stmt *ast.BranchStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitCaseClause(stmt *ast.CaseClause) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitCommClause(stmt *ast.CommClause) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitDeclStmt(stmt *ast.DeclStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitDeferStmt(stmt *ast.DeferStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitEmptyStmt(stmt *ast.EmptyStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitExprStmt(stmt *ast.ExprStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitForStmt(stmt *ast.ForStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitGoStmt(stmt *ast.GoStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitIfStmt(stmt *ast.IfStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitIncDecStmt(stmt *ast.IncDecStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitLabeledStmt(stmt *ast.LabeledStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitRangeStmt(stmt *ast.RangeStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitReturnStmt(stmt *ast.ReturnStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitSelectStmt(stmt *ast.SelectStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitSendStmt(stmt *ast.SendStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitSwitchStmt(stmt *ast.SwitchStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitTypeSwitchStmt(stmt *ast.TypeSwitchStmt) hlir.Operator {
	op := makeGoOperator(stmt)
	return op
}
