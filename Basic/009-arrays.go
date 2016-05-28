// Created: 2015-03-04, By Gavin
// Modified: 2015-03-04, By Gavin
// Link: https://tour.golang.org/moretypes/6

package main

import "fmt"

func main() {
	var array1 [2]string // 定义字符串数组
	array1[0] = "hello"
	array1[1] = "Gavin"

	fmt.Println(array1[0], array1[1])
	// Result: hello Gavin

	fmt.Println(array1)
	// Result: [hello Gavin]
}
