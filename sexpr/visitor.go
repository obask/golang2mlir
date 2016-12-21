package sexpr

import (
	"fmt"
	"reflect"
	"go/ast"
	"go/token"
)


type PrintASTVisitor struct {}

func (v *PrintASTVisitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return v
	}
	switch tt := n.(type) {
	case *ast.Package:
		break
	case *ast.File:
		break
	case *ast.GenDecl:
		if tt.Tok == token.TYPE {
			break
		}
	case *ast.TypeSpec:
		fmt.Println(tt.Name.Name)
		return nil
	default:
		fmt.Println(reflect.TypeOf(tt))
	}
	return v
}

