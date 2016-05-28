// Created: 2015-02-09, By Gavin
// Modified: 2015-02-09, By Gavin
// Link: http://tour.golang.org/flowcontrol/1

package main

import "fmt"

func main() {
	// Example 1
	sum := 0
	// for循环
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("从0加到9的结果是：", sum)

	// Example 2
	sum2 := 1
	// for 循环
	for j := 1; sum2 < 1000; {
		fmt.Println(j, sum2)
		sum2 += sum2
		j++
	}
}
