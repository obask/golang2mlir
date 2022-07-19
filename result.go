package main

import (
	"awesomeProject/sexpr"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main4() {

	filePath := "/Users/oleg/IdeaProjects/untitled/print_ast.go"

	fset := token.NewFileSet()

	_, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		println("Err")
		panic(err)
	}

	code := &ast.File{
		Doc:     nil,
		Package: 1,
		Name:    ast.NewIdent("main"),
		Decls: []ast.Decl{
			&ast.GenDecl{
				Doc:    nil,
				TokPos: 15,
				Tok:    token.IMPORT,
				Lparen: 22,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Doc:     nil,
						Name:    nil,
						Path:    sexpr.NewBasicLit("go/token"),
						Comment: nil,
						EndPos:  0,
					},
					&ast.ImportSpec{
						Doc:     nil,
						Name:    nil,
						Path:    sexpr.NewBasicLit("reflect"),
						Comment: nil,
						EndPos:  0,
					},
					&ast.ImportSpec{
						Doc:     nil,
						Name:    nil,
						Path:    sexpr.NewBasicLit("go/parser"),
						Comment: nil,
						EndPos:  0,
					},
					&ast.ImportSpec{
						Doc:     nil,
						Name:    nil,
						Path:    sexpr.NewBasicLit("./sexpr"),
						Comment: nil,
						EndPos:  0,
					},
					&ast.ImportSpec{
						Doc:     nil,
						Name:    nil,
						Path:    sexpr.NewBasicLit("fmt"),
						Comment: nil,
						EndPos:  0,
					},
				},
				Rparen: 78,
			},
			&ast.FuncDecl{
				Doc:  nil,
				Recv: nil,
				Name: ast.NewIdent("main"),
				Type: &ast.FuncType{
					Func: 141,
					Params: &ast.FieldList{
						Opening: 150,
						List:    nil,
						Closing: 151,
					},
					Results: nil,
				},
				Body: &ast.BlockStmt{
					Lbrace: 153,
					List: []ast.Stmt{
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("filePath"),
							},
							TokPos: 328,
							Tok:    47,
							Rhs: []ast.Expr{
								sexpr.NewBasicLit("/Users/oleg/IdeaProjects/untitled/print_ast.go"),
							},
						},
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("fset"),
							},
							TokPos: 387,
							Tok:    47,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("token"),
										Sel: ast.NewIdent("NewFileSet"),
									},
									Lparen:   406,
									Args:     nil,
									Ellipsis: 0,
									Rparen:   407,
								},
							},
						},
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("code"),
								ast.NewIdent("err"),
							},
							TokPos: 421,
							Tok:    47,
							Rhs: []ast.Expr{
								&ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   ast.NewIdent("parser"),
										Sel: ast.NewIdent("ParseFile"),
									},
									Lparen: 440,
									Args: []ast.Expr{
										ast.NewIdent("fset"),
										ast.NewIdent("filePath"),
										ast.NewIdent("nil"),
										sexpr.NewBasicLit(0),
									},
									Ellipsis: 0,
									Rparen:   463,
								},
							},
						},
						&ast.IfStmt{
							If:   466,
							Init: nil,
							Cond: &ast.BinaryExpr{
								X:     ast.NewIdent("err"),
								OpPos: 473,
								Op:    44,
								Y:     ast.NewIdent("nil"),
							},
							Body: &ast.BlockStmt{
								Lbrace: 480,
								List: []ast.Stmt{
									&ast.ExprStmt{
										X: &ast.CallExpr{
											Fun:    ast.NewIdent("println"),
											Lparen: 491,
											Args: []ast.Expr{
												sexpr.NewBasicLit("Err"),
											},
											Ellipsis: 0,
											Rparen:   497,
										},
									},
									&ast.ExprStmt{
										X: &ast.CallExpr{
											Fun:    ast.NewIdent("panic"),
											Lparen: 506,
											Args: []ast.Expr{
												ast.NewIdent("err"),
											},
											Ellipsis: 0,
											Rparen:   510,
										},
									},
								},
								Rbrace: 513,
							},
							Else: nil,
						},
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("fmt"),
									Sel: ast.NewIdent("Println"),
								},
								Lparen: 596,
								Args: []ast.Expr{
									sexpr.NewBasicLit("------------------------"),
								},
								Ellipsis: 0,
								Rparen:   623,
							},
						},
						&ast.AssignStmt{
							Lhs: []ast.Expr{
								ast.NewIdent("p"),
							},
							TokPos: 629,
							Tok:    47,
							Rhs: []ast.Expr{
								&ast.UnaryExpr{
									OpPos: 632,
									Op:    17,
									X: &ast.CompositeLit{
										Type: &ast.SelectorExpr{
											X:   ast.NewIdent("sexpr"),
											Sel: ast.NewIdent("SPrinter"),
										},
										Lbrace: 647,
										Elts:   nil,
										Rbrace: 648,
									},
								},
							},
						},
						&ast.ExprStmt{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{
									X:   ast.NewIdent("p"),
									Sel: ast.NewIdent("Sprint"),
								},
								Lparen: 659,
								Args: []ast.Expr{
									&ast.CallExpr{
										Fun: &ast.SelectorExpr{
											X:   ast.NewIdent("reflect"),
											Sel: ast.NewIdent("ValueOf"),
										},
										Lparen: 675,
										Args: []ast.Expr{
											ast.NewIdent("code"),
										},
										Ellipsis: 0,
										Rparen:   680,
									},
								},
								Ellipsis: 0,
								Rparen:   681,
							},
						},
					},
					Rbrace: 684,
				},
			},
		},
		Scope: nil,
		Imports: []*ast.ImportSpec{
			&ast.ImportSpec{
				Doc:     nil,
				Name:    nil,
				Path:    sexpr.NewBasicLit("go/token"),
				Comment: nil,
				EndPos:  0,
			},
			&ast.ImportSpec{
				Doc:     nil,
				Name:    nil,
				Path:    sexpr.NewBasicLit("reflect"),
				Comment: nil,
				EndPos:  0,
			},
			&ast.ImportSpec{
				Doc:     nil,
				Name:    nil,
				Path:    sexpr.NewBasicLit("go/parser"),
				Comment: nil,
				EndPos:  0,
			},
			&ast.ImportSpec{
				Doc:     nil,
				Name:    nil,
				Path:    sexpr.NewBasicLit("./sexpr"),
				Comment: nil,
				EndPos:  0,
			},
			&ast.ImportSpec{
				Doc:     nil,
				Name:    nil,
				Path:    sexpr.NewBasicLit("fmt"),
				Comment: nil,
				EndPos:  0,
			},
		},
		Unresolved: []*ast.Ident{
			ast.NewIdent("token"),
			ast.NewIdent("parser"),
			ast.NewIdent("nil"),
			ast.NewIdent("nil"),
			ast.NewIdent("println"),
			ast.NewIdent("panic"),
			ast.NewIdent("fmt"),
			ast.NewIdent("sexpr"),
			ast.NewIdent("reflect"),
		},
		Comments: nil,
	}

	//ast.Print(nil, code)
	printer.Fprint(os.Stdout, fset, code)

}
