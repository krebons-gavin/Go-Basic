// Created: 2015-03-04, By Gavin
// Modified: 2015-03-15, By Gavin
// Link: https://tour.golang.org/moretypes/7
// Link: https://tour.golang.org/moretypes/9

package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] = ", p[1:4])

	// 省略最小索引以为着从0开始
	fmt.Println("p[:3] ==", p[:3])

	// 省略最大索引意味着最大
	fmt.Println("p[4:] == ", p[4:])

	// 循环输出数组的所有元素
	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d \n", i, p[i])
	}

	// Making Slice
	// https://tour.golang.org/moretypes/9
	a := make([]int, 5)
	printSlice("a", a)
	// 定义长度为5的整形数组，默认值为0
	// a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b)
	// b len=0 cap=5 []

	c := b[:2]
	printSlice("c", c)
	// c len=2 cap=5 [0 0]

	d := c[2:5]
	printSlice("d", d)
	// d len=3 cap=3 [0 0 0]

	// New Example: nil-slices
	// Link: https://tour.golang.org/moretypes/10
	var z []int
	fmt.Println(z, len(z), cap(z))
	// [] 0 0
	if z == nil {
		fmt.Println("nil!")
		// nil!
		// nil 切片的值和容量都为0
	}
}

// 输出数组名称、长度、容量、数组本身
func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
