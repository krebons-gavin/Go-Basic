// Created: 2015-02-14, By Gavin
// Modified: 2015-02-14, By Gavin
// Link: https://tour.golang.org/flowcontrol/12

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	// 最后输出

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		// 逆序输出0到9
	}

	fmt.Println("done")
	// 先于world输出
}
