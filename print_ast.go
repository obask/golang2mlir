package main

import (
	"awesomeProject/hlir"
	"container/list"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"
	//"reflect"
	"go/parser"
)

func ololo() {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

//const np token.Pos = 0

//goland:noinspection GoUnhandledErrorResult
func main7() {

	println("----")

	op2 := hlir.Operator{
		Name:        "func",
		Dialect:     "go",
		Blocks:      nil,
		ReturnNames: nil,
		Attributes:  nil,
	}
	label := &hlir.Label{
		Name:        "^bb0",
		ParamValues: nil,
		ParamTypes:  nil,
	}
	bb0 := hlir.Block{
		Label: label,
		Items: []hlir.Operator{op2, op2},
	}

	op := hlir.Operator{
		Name:        "func",
		Dialect:     "go",
		Blocks:      []hlir.Block{bb0},
		ReturnNames: []hlir.ValueId{"%078"},
		Attributes:  nil,
	}

	op.RenderTo(os.Stdout, "  ")

	return

	// OLD version https://github.com/obask/homoiconic-go/blob/master/result.go
	filePath := "./hello.go"
	//filePath := "/Users/baskakov/IdeaProjects/homoiconic-go/assets/types.go"

	fset := token.NewFileSet()
	code, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		panic(err)
	}

	//ast.Print(fset, code)

	if err := printer.Fprint(os.Stdout, fset, code); err != nil {
		panic(err)
	}

	//p := &sexpr.SExpr{}
	//p.Sprint(reflect.ValueOf(code))
	println("---------")

	//_ = ast.Fprint(os.Stdout, fset, code, ast.NotNilFilter)

	Fprint(os.Stdout, fset, code, ast.NotNilFilter)

	tmp := &ast.BasicLit{
		ValuePos: 0,
		Kind:     token.STRING,
		Value:    "\"hello world\"",
	}
	printer.Fprint(os.Stdout, fset, tmp)

}
