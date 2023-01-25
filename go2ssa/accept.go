package go2ssa

import (
	"awesomeProject/ir"
	"go/ast"
)

func (c *GhostConvertor) visitDeclarations(items []ast.Decl) []ir.Operator {
	var result []ir.Operator
	for _, decl := range items {
		result = append(result, c.visitDecl(decl))
	}
	return result
}

func (c *GhostConvertor) visitDecl(decl ast.Decl) ir.Operator {
	switch node := decl.(type) {
	case *ast.BadDecl:
		return c.visitBadDecl(node)
	case *ast.GenDecl:
		return c.visitGenDecl(node)
	case *ast.FuncDecl:
		return c.visitFuncDecl(node)
	default:
		panic(node)
	}
}

func (c *GhostConvertor) visitExpressions(items []ast.Expr) []ir.Operator {
	var result []ir.Operator
	for _, expr := range items {
		result = append(result, c.visitExpr(expr))
	}
	return result
}

func (c *GhostConvertor) visitExpr(expr ast.Expr) ir.Operator {
	switch node := expr.(type) {
	case *ast.ArrayType:
		return c.visitArrayType(node)
	case *ast.BadExpr:
		return c.visitBadExpr(node)
	case *ast.BasicLit:
		return c.visitBasicLit(node)
	case *ast.BinaryExpr:
		return c.visitBinaryExpr(node)
	case *ast.CallExpr:
		return c.visitCallExpr(node)
	case *ast.ChanType:
		return c.visitChanType(node)
	case *ast.CompositeLit:
		return c.visitCompositeLit(node)
	case *ast.Ellipsis:
		return c.visitEllipsis(node)
	case *ast.FuncLit:
		return c.visitFuncLit(node)
	case *ast.FuncType:
		return c.visitFuncType(node)
	case *ast.Ident:
		return c.visitIdent(node)
	case *ast.IndexExpr:
		return c.visitIndexExpr(node)
	case *ast.IndexListExpr:
		return c.visitIndexListExpr(node)
	case *ast.InterfaceType:
		return c.visitInterfaceType(node)
	case *ast.KeyValueExpr:
		return c.visitKeyValueExpr(node)
	case *ast.MapType:
		return c.visitMapType(node)
	case *ast.ParenExpr:
		return c.visitParenExpr(node)
	case *ast.SelectorExpr:
		return c.visitSelectorExpr(node)
	case *ast.SliceExpr:
		return c.visitSliceExpr(node)
	case *ast.StarExpr:
		return c.visitStarExpr(node)
	case *ast.StructType:
		return c.visitStructType(node)
	case *ast.TypeAssertExpr:
		return c.visitTypeAssertExpr(node)
	case *ast.UnaryExpr:
		return c.visitUnaryExpr(node)
	default:
		panic(node)
	}
}

func (c *GhostConvertor) visitStatements(items []ast.Stmt) []ir.Operator {
	var result []ir.Operator
	for _, stmt := range items {
		result = append(result, c.visitStmt(stmt))
	}
	return result
}

func (c *GhostConvertor) visitStmt(stmt ast.Stmt) ir.Operator {
	switch node := stmt.(type) {
	case *ast.AssignStmt:
		return c.visitAssignStmt(node)
	case *ast.BadStmt:
		return c.visitBadStmt(node)
	case *ast.BlockStmt:
		return c.visitBlockStmt(node)
	case *ast.BranchStmt:
		return c.visitBranchStmt(node)
	case *ast.CaseClause:
		return c.visitCaseClause(node)
	case *ast.CommClause:
		return c.visitCommClause(node)
	case *ast.DeclStmt:
		return c.visitDeclStmt(node)
	case *ast.DeferStmt:
		return c.visitDeferStmt(node)
	case *ast.EmptyStmt:
		return c.visitEmptyStmt(node)
	case *ast.ExprStmt:
		return c.visitExprStmt(node)
	case *ast.ForStmt:
		return c.visitForStmt(node)
	case *ast.GoStmt:
		return c.visitGoStmt(node)
	case *ast.IfStmt:
		return c.visitIfStmt(node)
	case *ast.IncDecStmt:
		return c.visitIncDecStmt(node)
	case *ast.LabeledStmt:
		return c.visitLabeledStmt(node)
	case *ast.RangeStmt:
		return c.visitRangeStmt(node)
	case *ast.ReturnStmt:
		return c.visitReturnStmt(node)
	case *ast.SelectStmt:
		return c.visitSelectStmt(node)
	case *ast.SendStmt:
		return c.visitSendStmt(node)
	case *ast.SwitchStmt:
		return c.visitSwitchStmt(node)
	case *ast.TypeSwitchStmt:
		return c.visitTypeSwitchStmt(node)
	default:
		panic(node)
	}
}
