package sexpr

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unicode"
	"os"
	"strings"
	"go/ast"
	"go/token"
)


func NewBasicLit(value interface{}) *ast.BasicLit {
	switch value.(type) {
	case string:
		return &ast.BasicLit{ValuePos: token.NoPos, Kind: token.STRING, Value: fmt.Sprintf("%q", value)}
	case int:
		return &ast.BasicLit{ValuePos: token.NoPos, Kind: token.INT, Value: fmt.Sprint(value)}
	default:
		panic("NewIdent")
	}
}

type SPrinter struct {
	indent int                 // current indentation level
}

func IsExported(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(ch)
}

// NotNilFilter returns true for field values that are not nil;
// it returns false otherwise.
func NotNilFilter(_ string, v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return !v.IsNil()
	}
	return true
}

func IsBadField(v reflect.Value) bool {
	switch v.Interface().(type) {
	case *ast.CommentGroup:
		return true
	case token.Pos:
		return true

	default:
		return false
	}
}

// printf is a convenience wrapper that takes care of print errors.
func (p *SPrinter) printf(format string, args ...interface{}) {
	if _, err := fmt.Printf(format, args...); err != nil {
		panic(err)
	}
	if strings.Contains(format, "\n") {
		for j := p.indent; j > 0; j-- {
			os.Stdout.Write(indent)
			//fmt.Printf(indent)
		}
	}
}


// Implementation note: Print is written for AST nodes but could be
// used to print arbitrary data structures; such a version should
// probably be in a different package.
//
// Note: This code detects (some) cycles created via pointers but
// not cycles that are created via slices or maps containing the
// same slice or map. Code for general data structures probably
// should catch those as well.

func (p *SPrinter) Sprint(x reflect.Value) {
	if !ast.NotNilFilter("", x) {
		p.printf("nil")
		return
	}

	switch x.Kind() {
	case reflect.Interface:
		p.Sprint(x.Elem())

	case reflect.Map:
		panic(nil)

	case reflect.Ptr:
		// type-checked ASTs may contain cycles - use ptrmap
		// to keep track of objects that have been printed
		// already and print the respective line number instead
		elem := x.Elem()
		if elem.Kind() == reflect.Struct {
			switch elem.Interface().(type) {
			case ast.Ident:
				name := elem.FieldByName("Name").String()
				p.printf("ast.NewIdent(%q)", name)
				return
			case ast.BasicLit:
				value := elem.FieldByName("Value").String()
				p.printf("sexpr.NewBasicLit(%v)", value)
				return
			case ast.Scope:
				p.printf("nil")
				return
			}
		}
		p.printf("&")
		p.Sprint(elem)


	case reflect.Array:
		panic(nil)

	case reflect.Slice:
		p.printf("%s {", x.Type())
		if x.Len() > 0 {
			p.indent++
			p.printf("\n")
			for i, n := 0, x.Len(); i < n; i++ {
				p.Sprint(x.Index(i))
				p.printf(",\n")
			}
			p.indent--
		}
		p.printf("}")

	case reflect.Struct:
		t := x.Type()
		p.printf("%s {", t)
		p.indent++
		first := true
		for i, n := 0, t.NumField(); i < n; i++ {
			// exclude non-exported fields because their
			// values cannot be accessed via reflection
			if name := t.Field(i).Name; IsExported(name) {
				value := x.Field(i)
				if (IsBadField(value)) {
					continue
				}
				if first {
					p.printf("\n")
					first = false
				}
				p.printf("%s: ", name)
				p.Sprint(value)
				p.printf(",\n")
			}
		}
		p.indent--
		p.printf("}")

	default:
		v := x.Interface()
		switch v.(type) {
		case string:
			// print strings in quotes
			p.printf("%q", v)
		case int:
			p.printf("%d", v)
		case token.Pos:
			p.printf("%v", v)
		case token.Token:
			i2 := v.(token.Token)
			if i2.IsKeyword() {
				p.printf("token.%s", strings.ToUpper(fmt.Sprint(v)))
			} else if i2.IsOperator() {
				p.printf("%d", v)
			} else {
				panic(i2)
			}
		default:
			panic(reflect.TypeOf(v))
		}
	}
}


