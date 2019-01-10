package main

import "fmt"

var p = f()

func main() {
	fmt.Println(f() == f())
}

func f() *int {
	v := 1
	return &v
}
