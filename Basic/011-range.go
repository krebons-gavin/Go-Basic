// Created: 2015-03-24, By Gavin
// Modified: 2015-03-24, By Gavin
// Link: https://tour.golang.org/moretypes/12
// Link: https://tour.golang.org/moretypes/13

package main

import "fmt"

// 指数：2的n次幂
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Example2: https://tour.golang.org/moretypes/13
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i)
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}
