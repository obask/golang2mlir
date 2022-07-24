// Go has strange behavior -- in one case it doesn't support generics
// In over case it can't convert []AnyType to []interface[] even pointers
// So it don't support sub-typing that makes generic coding really hard

package main

import (
	"awesomeProject/go2ssa"
	"awesomeProject/sexpr"
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"io/ioutil"
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

	return
}

//return
//op2 := mlir.Operator{
//	Name:       "func",
//	Dialect:    "go",
//	Regions:    nil,
//	ReturnName: "",
//	Attributes: map[string]mlir.Attribute{"dfadsfg": mlir.StringAttr("\"dsad\""), "dfadsfg2": mlir.NumberAttr(123)},
//}
//label := &mlir.BlockLabel{
//	Name:        "^bb0",
//	ParamValues: nil,
//	ParamTypes:  nil,
//}
//bb0 := mlir.BasicBlock{
//	Label: label,
//	Items: []mlir.Operator{op2, op2},
//}
//
//op := mlir.Operator{
//	Name:       "func",
//	Dialect:    "go",
//	Regions:    []mlir.Region{[]mlir.BasicBlock{bb0, bb0}},
//	ReturnName: "%078",
//	Attributes: map[string]mlir.Attribute{"symbol_name": mlir.StringAttr("@main")},
//}

func main1() {
	//filename := "./assets/code.clj"
	filename := "/Users/baskakov/IdeaProjects/homoiconic-go/assets/types.go"
	tree := sexpr.ParseFile(token.NewFileSet(), filename)

	fmt.Println("dbg 3")

	//	tree := ABranch{val: []ATree{
	//		ASymbol{val: "main"},
	//		ASymbol{val: "dsa"},
	//		ABranch{val: []ATree{
	//			ASymbol{val: "+"},
	//			ANumber{val: "1"},
	//			ANumber{val: "2"},
	//			},},
	//	},}

	fmt.Println(tree)

}

func main2() {

	filename := "./resources/result.clj"
	src, err := ioutil.ReadFile(filename)
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

func process(node sexpr.ANode) reflect.Value {
	switch node.(type) {
	case sexpr.ABranch:
		name := node.Name()
		switch name {
		case "const":
			repr := node.Children()[0].String()
			return reflect.ValueOf(&ast.BasicLit{Kind: token.STRING, Value: repr})
		case "%":
			fields := []reflect.Value{}
			for _, child := range node.Children() {
				fields = append(fields, process(child))
			}
			return sexpr.CreateSlice(fields)
		}
		// default:
		fields := []reflect.Value{}
		for _, child := range node.Children() {
			fields = append(fields, process(child))
		}
		t, ok := sexpr.TypeIndex[name]
		if !ok {
			panic("type not found: " + name)
		}
		return sexpr.CreateStruct(t, fields)
	case sexpr.AString:
		ss := node.String()
		return reflect.ValueOf(ast.NewIdent(ss[1 : len(ss)-1]))

	case sexpr.ASymbol:
		val := node.String()
		switch val {
		case "nil":
			return reflect.ValueOf(nil)
		case "true":
			return reflect.ValueOf(true)
		case "false":
			return reflect.ValueOf(false)
		case "import":
			return reflect.ValueOf(token.IMPORT)
		case "type":
			return reflect.ValueOf(token.TYPE)
		default:
			panic(val)
		}
	}
	fmt.Println("process")
	fmt.Println("Children: ", node.Children())
	child := node.Children()[1]
	fmt.Println(child.Name())

	panic(nil)
}

func main3() {

	filename2 := "./assets/types.go"
	tree2, _ := parser.ParseFile(token.NewFileSet(), filename2, nil, 0)

	filename := "./assets/code.clj"
	file := sexpr.ParseFile(token.NewFileSet(), filename)
	tree := file.(sexpr.ABranch).Val[0]
	//fmt.Println(tree)

	result := process(tree).Interface().(*ast.File)

	printer := &sexpr.SPrinter{}
	printer.Sprint(reflect.ValueOf(tree2.Decls[1]))
	fmt.Println()
	fmt.Println("--------------")
	printer.Sprint(reflect.ValueOf(result.Decls[1]))

	fmt.Println()
}
