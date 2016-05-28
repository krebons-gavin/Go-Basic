// Created: 2015-02-06, By Gavin
// Modified: 2015-02-06, By Gavin
// Link: http://tour.golang.org/basics/4

package main

import (
	"fmt"
	"math"
)

// 加法函数，返回两个数的和
// 注意：要加上变量及返回值的类型
func addxy(x int, y int) int {
	return (x + y)
}

// 字符串交换
// http://tour.golang.org/basics/6
func swap(x, y string) (string, string) {
	return y, x
}

// http://tour.golang.org/basics/7
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 定义Bool型变量, 未赋值的话系统默认为 false
var c, python, java bool

// 主函数
func main() {
	fmt.Println("Hello, Gavin, Today is 2015-02-05.")
	fmt.Println(math.Pi)       // 输出Pi
	fmt.Println(addxy(42, 13)) // 调用加法函数

	// 字符串交换
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println("The number 17 after split is:")
	fmt.Println(split(17))
	// 7 10

	fmt.Println("The number 39 after split is:")
	fmt.Println(split(39))
	// 17 22

	// 定义整形变量，未赋值的话系统默认为0
	var i int
	fmt.Println(i, c, python, java)
	// 0 false false false

	// 求平方根
	var num1, num2 int = 3, 4
	var f float64 = math.Sqrt(float64(num1*num1 + num2*num2))
	var z int = int(f) // 强制类型转换
	fmt.Println(num1, "与", num2, "的平方根是：", z)
}
