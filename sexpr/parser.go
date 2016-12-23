package sexpr

import (
	"fmt"
	"go/token"
	"io/ioutil"
	"go/scanner"
)

type ANode interface {
	String() string
	Name() string
	Children() []ANode
}

type ABranch struct {
	ANode
	Val []ANode // denoted object; or nil
}

type ALeaf interface {
	ANode
}

type ANumber struct {
	ALeaf
	Val  string // denoted object; or nil
	kind token.Token
}

type ASymbol struct {
	ALeaf
	Val string // denoted object; or nil
}

type AString struct {
	ALeaf
	Val string // denoted object; or nil
}

// String returns the string corresponding to the token tok.
// For operators, delimiters, and keywords the string is the actual
// token character sequence (e.g., for the token ADD, the string is
// "+"). For all other tokens the string corresponds to the token
// constant name (e.g. for the token IDENT, the string is "IDENT").
//
func (stmt ABranch) String() string {
	res := "["
	for _, x := range stmt.Val {
		res += x.String()
		res += " "
	}
	res += "]"
	return res
}

func (node ABranch) Name() string {
	return node.Val[0].(ASymbol).Val
}

func (node ABranch) Children() []ANode {
	return node.Val[1:]
}

func (stmt ANumber) String() string {
	return stmt.Val
}

func (stmt ASymbol) String() string {
	return stmt.Val
}

func (stmt AString) String() string {
	return stmt.Val
}

func parseCode(lexer *scanner.Scanner, state []ANode) ANode {
	_, tok, lit := lexer.Scan()
	// skip endlines
	for tok == token.SEMICOLON {
		_, tok, lit = lexer.Scan()
	}
	switch tok {
	case token.EOF:
		fmt.Println("dbg: token.EOF")
		return ABranch{Val: state}

	case token.LPAREN, token.LBRACK, token.LBRACE:
		fmt.Println("dbg: token.LPAREN")
		expr := parseCode(lexer, []ANode{})
		return parseCode(lexer, append(state, expr))

	case token.RPAREN, token.RBRACK, token.RBRACE:
		fmt.Println("dbg: token.RPAREN")
		return ABranch{Val: state}

	case token.INT, token.FLOAT:
		fmt.Println("dbg: ANumber")
		newElem := ANumber{Val: lit, kind: tok}
		return parseCode(lexer, append(state, newElem))

	case token.STRING, token.CHAR, token.CONST:
		fmt.Println("dbg: AString = " + lit)
		newElem := AString{Val: lit}
		return parseCode(lexer, append(state, newElem))
	}
	switch {
	case tok.IsOperator() || tok.IsKeyword():
		// IsOperator has a bug with brackets and semicolon!!
		fmt.Println("dbg: token.IsOperator = " + tok.String())
		newElem := ASymbol{Val: tok.String()}
		return parseCode(lexer, append(state, newElem))

	case tok == token.IDENT || tok ==  token.IF:
		// IsOperator has a bug with brackets and semicolon!!
		fmt.Println("dbg: token.IsOperator = " + lit)
		newElem := ASymbol{Val: lit}
		return parseCode(lexer, append(state, newElem))

	default:
		// Unknown token
		fmt.Println("dbg: default")
		fmt.Printf("\t%s    ->  %q\n", tok, lit)
		panic("SUCCESS default branch of parseCode")
		return nil
	}
}

//default:
////                    System.out.println("L_ATOM");
////                    System.out.println(state.getClass());
//state.add(new AString(curr));
//return reduceTree(tokens, state);
//}
//} else {
//return new ABranch(state);
//}

func ParseFile(fset *token.FileSet, filename string) ANode {

	fmt.Println("dbg 1")

	// get source
	text, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	var lexer scanner.Scanner

	fileSet := fset.AddFile(filename, -1, len(text))

	errorHandler := func(pos token.Position, msg string) {
		// FIXME this happened for ILLEGAL tokens, forex '?'
		panic("SUCCESS in scanner errorHandler")
	}

	var m scanner.Mode
	lexer.Init(fileSet, text, errorHandler, m)

	fmt.Println("dbg 2.1")

	// Repeated calls to Scan yield the token sequence found in the input.
	//	for {
	//		_, tok, lit := lexer.Scan()
	//		if tok == token.EOF {
	//			break
	//		}
	//		fmt.Printf("\t%s    %q\n", tok, lit)
	//	}
	//

	fmt.Println("dbg 2.2")

	return parseCode(&lexer, []ANode{})
}
