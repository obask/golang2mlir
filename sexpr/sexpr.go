package sexpr

import (
	"fmt"
	"reflect"
	"os"
	"strings"
	"go/ast"
	"go/token"
)

type SExpr struct {
	indent int                 // current indentation level
}


var indent = []byte("  ")

// printf is a convenience wrapper that takes care of print errors.
func (p *SExpr) printf(format string, args ...interface{}) {
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

func (p *SExpr) Sprint(x reflect.Value) {
	switch x.Kind() {
	case reflect.Interface, reflect.Ptr:
		if x.IsNil() {
			p.printf("nil")
			return
		}
		p.Sprint(x.Elem())

	case reflect.Slice:
		if x.IsNil() || x.Len() == 0 {
			p.printf("%s", "[% ]")
			return
		}
		p.printf("%s", "[% ")
		p.indent+=2
		for i, n := 0, x.Len(); i < n; i++ {
			p.Sprint(x.Index(i))
			p.printf("\n")
		}
		p.printf("]")
		p.indent-=2

	case reflect.Struct:
		switch x.Interface().(type) {
		case ast.Ident:
			name := x.FieldByName("Name").String()
			p.printf("(i %q)", name)
		case ast.BasicLit:
			value := x.FieldByName("Value").String()
			p.printf("(const %v)", value)
		case ast.Scope:
			p.printf("nil")
		default:
			t := x.Type()
			p.printf("(%s ", t.String()[4:])
			p.indent++
			first := true
			for i, n := 0, t.NumField(); i < n; i++ {
				// exclude non-exported fields because their
				// values cannot be accessed via reflection
				if name := t.Field(i).Name; IsExported(name) {
					value := x.Field(i)
					if (isBadField(value)) {
						continue
					}
					if first {
						p.printf("\n")
						first = false
					}
					//p.printf("%s: ", name)
					p.Sprint(value)
					p.printf("\n")
				}
			}
			p.indent--
			p.printf(")")
		}

	case reflect.String:
		p.printf("%q", x.Interface())

	default:
		// simple types
		v := x.Interface()
		switch v.(type) {
		case int:
			p.printf("%d", v)
		case token.Pos:
			p.printf("%v", v)
		case token.Token:
			i2 := v.(token.Token)
			if i2 <= token.STRING {
				panic(i2)
			} else if i2 <= token.COLON {
				p.printf("%q", v)
			} else {
				p.printf("token.%s", strings.ToUpper(fmt.Sprint(v)))
			}
		default:
			panic(reflect.TypeOf(v))
		}
	}
}


