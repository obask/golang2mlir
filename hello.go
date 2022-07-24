package main

import "fmt"

func hello() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
}
