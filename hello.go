package main

import "fmt"

func hello(i int) {
	if i <= 42 {
		fmt.Println(i)
		i = i + 1
	} else {
		fmt.Println("Hello")
	}
}

//type A struct {
//	Field1 string
//}
//
//type B struct {
//	A
//	Field2 string
//}
//
//fn makeB() *B {
//	return &B{
//		A{"das"},
//		"das",
//	}
//}
//
//fn test(i int) {
//	a := &A{"qwe"}
//	fmt.Println(a)
//	a = makeB()
//}
