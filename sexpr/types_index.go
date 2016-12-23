package sexpr

import (
	"go/ast"
	"reflect"
)

var types []interface{} = []interface{} {
	ast.Comment{},
	ast.CommentGroup{},
	ast.Field{},
	ast.FieldList{},
	ast.BadExpr{},
	ast.Ident{},
	ast.Ellipsis{},
	ast.BasicLit{},
	ast.FuncLit{},
	ast.CompositeLit{},
	ast.ParenExpr{},
	ast.SelectorExpr{},
	ast.IndexExpr{},
	ast.SliceExpr{},
	ast.TypeAssertExpr{},
	ast.CallExpr{},
	ast.StarExpr{},
	ast.UnaryExpr{},
	ast.BinaryExpr{},
	ast.KeyValueExpr{},
	ast.ArrayType{},
	ast.StructType{},
	ast.FuncType{},
	ast.InterfaceType{},
	ast.MapType{},
	ast.ChanType{},
	ast.BadStmt{},
	ast.DeclStmt{},
	ast.EmptyStmt{},
	ast.LabeledStmt{},
	ast.ExprStmt{},
	ast.SendStmt{},
	ast.IncDecStmt{},
	ast.AssignStmt{},
	ast.GoStmt{},
	ast.DeferStmt{},
	ast.ReturnStmt{},
	ast.BranchStmt{},
	ast.BlockStmt{},
	ast.IfStmt{},
	ast.CaseClause{},
	ast.SwitchStmt{},
	ast.TypeSwitchStmt{},
	ast.CommClause{},
	ast.SelectStmt{},
	ast.ForStmt{},
	ast.RangeStmt{},
	ast.ImportSpec{},
	ast.ValueSpec{},
	ast.TypeSpec{},
	ast.BadDecl{},
	ast.GenDecl{},
	ast.FuncDecl{},
	ast.File{},
	ast.Package{},
}

func translateTypeIndex(typesList []interface{}) map[string]reflect.Type {
	result := map[string]reflect.Type{}
	for _,x := range typesList {
		t := reflect.TypeOf(x)
		result[t.Name()] = t
	}
	return result
}

var TypeIndex map[string]reflect.Type = translateTypeIndex(types)
