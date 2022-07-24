package go2ssa

import (
	"awesomeProject/mlir"
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
	Result         *mlir.Operator
	pos            int
	insertionPoint [][]mlir.Operator
}

func makeGoOperator(node ast.Node) mlir.Operator {
	typeName := fmt.Sprint(reflect.TypeOf(node))[5:]
	return mlir.Operator{
		Name:    strings.ToLower(typeName[:1]) + typeName[1:],
		Dialect: "go",
	}
}

func (c *GhostConvertor) processRegion(op *mlir.Operator, items []mlir.Operator) {
	op.Regions = append(op.Regions, []mlir.BasicBlock{{Items: items}})
}

func (c *GhostConvertor) processOperands(m *mlir.Operator, args []ast.Expr) {
	c.insertionPoint = append(c.insertionPoint, nil)

	c.insertionPoint = c.insertionPoint[:len(c.insertionPoint)-1]
}

func (c *GhostConvertor) visitFile(file *ast.File) mlir.Operator {
	op := makeGoOperator(file)
	op.Attributes["Imports"] = mlir.StringAttr(fmt.Sprintf("%v", file.Imports))
	c.processRegion(&op, c.visitDeclarations(file.Decls))
	newVisitor := GhostVisitor{}
	for _, decl := range file.Decls {
		ast.Walk(&newVisitor, decl)
	}

	return op
}

func (c *GhostConvertor) visitBadDecl(decl *ast.BadDecl) mlir.Operator {
	op := makeGoOperator(decl)
	return op
}

func (c *GhostConvertor) visitGenDecl(decl *ast.GenDecl) mlir.Operator {
	op := makeGoOperator(decl)
	return op
}

func (c *GhostConvertor) visitFuncDecl(decl *ast.FuncDecl) mlir.Operator {
	op := makeGoOperator(decl)
	op.Attributes["Recv"] = mlir.StringAttr(fmt.Sprintf("%v", decl.Recv))
	op.Attributes["Name"] = mlir.StringAttr(fmt.Sprintf("%v", decl.Name))
	op.Attributes["Type"] = mlir.StringAttr(fmt.Sprintf("%v", decl.Type))
	c.processRegion(&op, c.visitStatements(decl.Body.List))
	return op
}

func (c *GhostConvertor) visitArrayType(expr *ast.ArrayType) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitBadExpr(expr *ast.BadExpr) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitBasicLit(expr *ast.BasicLit) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitBinaryExpr(expr *ast.BinaryExpr) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitCallExpr(expr *ast.CallExpr) mlir.Operator {
	op := makeGoOperator(expr)
	switch f := expr.Fun.(type) {
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
	c.processOperands(&op, expr.Args)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitChanType(expr *ast.ChanType) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitCompositeLit(expr *ast.CompositeLit) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitEllipsis(expr *ast.Ellipsis) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitFuncLit(expr *ast.FuncLit) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitFuncType(expr *ast.FuncType) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitIdent(expr *ast.Ident) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitIndexExpr(expr *ast.IndexExpr) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitIndexListExpr(expr *ast.IndexListExpr) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitInterfaceType(expr *ast.InterfaceType) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitKeyValueExpr(expr *ast.KeyValueExpr) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitMapType(expr *ast.MapType) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitParenExpr(expr *ast.ParenExpr) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitSelectorExpr(expr *ast.SelectorExpr) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitSliceExpr(expr *ast.SliceExpr) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitStarExpr(expr *ast.StarExpr) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitStructType(expr *ast.StructType) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitTypeAssertExpr(expr *ast.TypeAssertExpr) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitUnaryExpr(expr *ast.UnaryExpr) mlir.Operator {
	op := makeGoOperator(expr)
	op.SetReturnName()
	return op
}

func (c *GhostConvertor) visitAssignStmt(stmt *ast.AssignStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitBadStmt(stmt *ast.BadStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitBlockStmt(stmt *ast.BlockStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitBranchStmt(stmt *ast.BranchStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitCaseClause(stmt *ast.CaseClause) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitCommClause(stmt *ast.CommClause) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitDeclStmt(stmt *ast.DeclStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitDeferStmt(stmt *ast.DeferStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitEmptyStmt(stmt *ast.EmptyStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitExprStmt(stmt *ast.ExprStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitForStmt(stmt *ast.ForStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitGoStmt(stmt *ast.GoStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitIfStmt(stmt *ast.IfStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitIncDecStmt(stmt *ast.IncDecStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitLabeledStmt(stmt *ast.LabeledStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitRangeStmt(stmt *ast.RangeStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitReturnStmt(stmt *ast.ReturnStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitSelectStmt(stmt *ast.SelectStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitSendStmt(stmt *ast.SendStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitSwitchStmt(stmt *ast.SwitchStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}

func (c *GhostConvertor) visitTypeSwitchStmt(stmt *ast.TypeSwitchStmt) mlir.Operator {
	op := makeGoOperator(stmt)
	return op
}
