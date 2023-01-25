package go2ssa

import (
	"awesomeProject/ir"
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
	Result         *ir.Operator
	pos            int
	insertionPoint [][]ir.Operator
}

func makeGoOperator(node ast.Node) ir.Operator {
	typeName := fmt.Sprint(reflect.TypeOf(node))[5:]
	return ir.Operator{
		Name:    strings.ToLower(typeName[:1]) + typeName[1:],
		Dialect: "go",
	}
}

func (c *GhostConvertor) processRegion(op *ir.Operator, items []ir.Operator) {
	op.Blocks = append(op.Blocks, ir.Block{Items: items})
}

func (c *GhostConvertor) processOperands(m *ir.Operator, args []ast.Expr) {
	c.insertionPoint = append(c.insertionPoint, nil)

	c.insertionPoint = c.insertionPoint[:len(c.insertionPoint)-1]
}

func (c *GhostConvertor) visitFile(file *ast.File) ir.Operator {
	op := makeGoOperator(file)
	op.Attributes["Imports"] = ir.StringAttr(fmt.Sprintf("%v", file.Imports))
	c.processRegion(&op, c.visitDeclarations(file.Decls))
	newVisitor := GhostVisitor{}
	for _, decl := range file.Decls {
		ast.Walk(&newVisitor, decl)
	}

	return op
}

func (c *GhostConvertor) visitBadDecl(decl *ast.BadDecl) ir.Operator {
	op := makeGoOperator(decl)
	return op
}

func (c *GhostConvertor) visitGenDecl(decl *ast.GenDecl) ir.Operator {
	op := makeGoOperator(decl)
	return op
}

func (c *GhostConvertor) visitFuncDecl(decl *ast.FuncDecl) ir.Operator {
	op := makeGoOperator(decl)
	op.Attributes["Recv"] = ir.StringAttr(fmt.Sprintf("%v", decl.Recv))
	op.Attributes["Name"] = ir.StringAttr(fmt.Sprintf("%v", decl.Name))
	op.Attributes["Type"] = ir.StringAttr(fmt.Sprintf("%v", decl.Type))
	c.processRegion(&op, c.visitStatements(decl.Body.List))
	return op
}

func (c *GhostConvertor) visitArrayType(expr *ast.ArrayType) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitBadExpr(expr *ast.BadExpr) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitBasicLit(expr *ast.BasicLit) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitBinaryExpr(expr *ast.BinaryExpr) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitCallExpr(expr *ast.CallExpr) ir.Operator {
	op := makeGoOperator(expr)
	switch f := expr.Fun.(type) {
	case *ast.Ident:
		op.Attributes["name"] = ir.StringAttr(f.Name)
	case *ast.SelectorExpr:
		switch lhs := f.X.(type) {
		case *ast.Ident:
			op.Attributes["name"] = ir.StringAttr(lhs.Name + "." + f.Sel.Name)
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

func (c *GhostConvertor) visitChanType(expr *ast.ChanType) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitCompositeLit(expr *ast.CompositeLit) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitEllipsis(expr *ast.Ellipsis) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitFuncLit(expr *ast.FuncLit) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitFuncType(expr *ast.FuncType) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitIdent(expr *ast.Ident) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitIndexExpr(expr *ast.IndexExpr) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitIndexListExpr(expr *ast.IndexListExpr) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitInterfaceType(expr *ast.InterfaceType) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitKeyValueExpr(expr *ast.KeyValueExpr) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitMapType(expr *ast.MapType) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitParenExpr(expr *ast.ParenExpr) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitSelectorExpr(expr *ast.SelectorExpr) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitSliceExpr(expr *ast.SliceExpr) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitStarExpr(expr *ast.StarExpr) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitStructType(expr *ast.StructType) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitTypeAssertExpr(expr *ast.TypeAssertExpr) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitUnaryExpr(expr *ast.UnaryExpr) ir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitAssignStmt(stmt *ast.AssignStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitBadStmt(stmt *ast.BadStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitBlockStmt(stmt *ast.BlockStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitBranchStmt(stmt *ast.BranchStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitCaseClause(stmt *ast.CaseClause) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitCommClause(stmt *ast.CommClause) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitDeclStmt(stmt *ast.DeclStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitDeferStmt(stmt *ast.DeferStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitEmptyStmt(stmt *ast.EmptyStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitExprStmt(stmt *ast.ExprStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitForStmt(stmt *ast.ForStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitGoStmt(stmt *ast.GoStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitIfStmt(stmt *ast.IfStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitIncDecStmt(stmt *ast.IncDecStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitLabeledStmt(stmt *ast.LabeledStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitRangeStmt(stmt *ast.RangeStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitReturnStmt(stmt *ast.ReturnStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitSelectStmt(stmt *ast.SelectStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitSendStmt(stmt *ast.SendStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitSwitchStmt(stmt *ast.SwitchStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitTypeSwitchStmt(stmt *ast.TypeSwitchStmt) ir.Operator {
	op := makeGoOperator(stmt)
	return op
}
