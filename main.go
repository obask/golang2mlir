// Go has strange behavior -- in one case it doesn't support generics
// In over case it can't convert []AnyType to []interface[] even pointers
// So it doesn't support sub-typing that makes generic coding really hard

package main

import (
	"awesomeProject/go2ssa"
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"os"
	"reflect"
)

func _fake() {
	ast.Print(nil, nil)
	token.NewFileSet()
	reflect.ValueOf(nil)
	fmt.Print()
	parser.ParseExpr("")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GoydaFilter(k string, v reflect.Value) bool {
	if v.Type().AssignableTo(reflect.TypeOf(token.Pos(0))) {
		return false
	}
	switch k {
	case "Obj":
		return false
	}
	return ast.NotNilFilter(k, v)
}

func main() {
	println("----")

	filePath := "./hello.go"
	fset := token.NewFileSet()
	code, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		panic(err)
	}

	g := &go2ssa.GhostVisitor{}

	ast.Walk(g, code)

	//fmt.Printf("%+v\n", g.Result)
	println("render:")
	g.Result.RenderTo(os.Stdout, "")

	Fprint(os.Stdout, fset, code, GoydaFilter)

	return
}

//return
//op2 := hlir.Operator{
//	Name:       "func",
//	Dialect:    "go",
//	Blocks:    nil,
//	ReturnNames: "",
//	Attributes: map[string]hlir.Attribute{"dfadsfg": hlir.StringAttr("\"dsad\""), "dfadsfg2": hlir.NumberAttr(123)},
//}
//label := &hlir.BlockLabel{
//	Name:        "^bb0",
//	ParamValues: nil,
//	ParamTypes:  nil,
//}
//bb0 := hlir.BasicBlock{
//	Label: label,
//	Items: []hlir.Operator{op2, op2},
//}
//
//op := hlir.Operator{
//	Name:       "func",
//	Dialect:    "go",
//	Blocks:    []hlir.Region{[]hlir.BasicBlock{bb0, bb0}},
//	ReturnNames: "%078",
//	Attributes: map[string]hlir.Attribute{"symbol_name": hlir.StringAttr("@main")},
//}

func main2() {

	filename := "./resources/result.clj"
	src, err := os.ReadFile(filename)
	check(err)

	//	src := []byte("cos(x) + 1i*sin(x) // Euler")

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		_, tok, lit := s.Scan()
		if tok == token.SEMICOLON {
			continue
		}
		if tok == token.EOF {
			break
		}
		fmt.Printf("\t%s    %q\n", tok, lit)
		//		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}

}
