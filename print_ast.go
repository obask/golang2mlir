package main

import (
	"go/token"
	"reflect"
	"go/parser"
	"./sexpr"
)

// TODO update Visitor pattern to visit and un-visit nodes

func main() {

	//filePath := "./print_ast.go"
	filePath := "/Users/baskakov/IdeaProjects/homoiconic-go/assets/types.go"

	code, err := parser.ParseFile(token.NewFileSet(), filePath, nil, 0)
	if err != nil {
		panic(err)
	}

	//ast.Print(fset, code)

	//printer.Fprint(os.Stdout, fset, code)

	p := &sexpr.SExpr{}
	p.Sprint(reflect.ValueOf(code))
	println()

}
