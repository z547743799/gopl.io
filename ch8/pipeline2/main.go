// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 229.

// Pipeline2 demonstrates a finite 3-stage pipeline.
package main

import "fmt"

//!+
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	//所有数据全部发送成功时才需要关闭channle

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

//!-
