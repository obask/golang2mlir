package main

import (
	"go/token"
	"reflect"
	"go/parser"
	"./sexpr"
	"fmt"
)

// TODO update Visitor pattern to visit and un-visit nodes

func main() {
	//fset := token.NewFileSet()
	//p, _ := ioutil.ReadFile()
	//file, _ := parser.ParseFile(fset, "print_ast.go", p, parser.ParseComments)
	//ast.Print(fset, file)

	filePath := "/Users/oleg/IdeaProjects/untitled/print_ast.go"

	fset := token.NewFileSet()

	code, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		println("Err")
		panic(err)
	}

	//ast.Print(fset, code)

	//printer.Fprint(os.Stdout, fset, code)

	fmt.Println("------------------------")

	p := &sexpr.SPrinter{}
	p.Sprint(reflect.ValueOf(code))

}
