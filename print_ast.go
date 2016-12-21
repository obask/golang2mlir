package main

import (
	"go/token"
	"reflect"
	"go/parser"
	"./sexpr"
)

// TODO update Visitor pattern to visit and un-visit nodes

func main() {

	filePath := "/Users/oleg/IdeaProjects/untitled/print_ast.go"
	fset := token.NewFileSet()

	code, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		println("Err")
		panic(err)
	}

	//ast.Print(fset, code)

	//printer.Fprint(os.Stdout, fset, code)

	p := &sexpr.SExpr{}
	p.Sprint(reflect.ValueOf(code))
	println()

}
