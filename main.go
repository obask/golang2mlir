// Go has strange behavior -- in one case it doesn't support generics
// In over case it can't convert []AnyType to []interface[] even pointers
// So it don't support sub-typing that makes generic coding really hard

package main

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
	"./sexpr"
	"go/ast"
	"reflect"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

func filter(node sexpr.ANode) []sexpr.ANode {
	println("filer")
	fmt.Println("Name: ", node.Name())
	fmt.Println("Children: ", node.Children())
	child := node.Children()[1]
	fmt.Println(child.Name())
	return []sexpr.ANode{}
}

func _fake() {
	ast.Print(nil, nil)
	token.NewFileSet()
}

type Test struct {
		       //Doc     *ast.CommentGroup // associated documentation; or nil
	Names ast.Expr // value names (len(Names) > 0)
		       //Type    ast.Expr          // value type; or nil
		       //Values  []ast.Expr        // initial values; or nil
		       //Comment *ast.CommentGroup // line comments; or nil
}

//type ValueSpec struct {
//	Doc     *CommentGroup // associated documentation; or nil
//	Names   []*Ident      // value names (len(Names) > 0)
//	Type    Expr          // value type; or nil
//	Values  []Expr        // initial values; or nil
//	Comment *CommentGroup // line comments; or nil
//}

func createStruct(t reflect.Type, fields []reflect.Value) reflect.Value {
	item := reflect.New(t).Elem()
	pos := 0
	n := t.NumField()
	for i := 0; i < n; i++ {
		// exclude non-exported fields because their
		// values cannot be accessed via reflection
		curr := item.Field(i)
		if (!sexpr.IsBadField(curr)) {
			curr.Set(fields[pos])
			pos++
		}
	}
	return item
}

func main() {
	//filename := "/Users/baskakov/IdeaProjects/homoiconic-go/assets/code.clj"
	//file := sexpr.ParseFile(token.NewFileSet(), filename)
	//tree := file.(sexpr.ABranch).Val[0]
	//filter(tree)

	//fmt.Println(tree)
	//ast.Print(token.NewFileSet(), tree)

	//fmt.Println(tree)

	//tree.(sexpr.ABranch)

	idents := []*ast.Ident{ast.NewIdent("xxx")}
	of1 := reflect.ValueOf(idents)
	of2 := reflect.ValueOf(ast.CallExpr{}).Addr()
	exps := []ast.Expr{&ast.CallExpr{}}

	fields := []reflect.Value{of1, of2, reflect.ValueOf(exps)}
	value := createStruct(reflect.TypeOf(ast.ValueSpec{}), fields)

	ast.Print(token.NewFileSet(), value.Interface())

	//fmt.Println("value =", value)

}



