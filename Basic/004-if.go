// Created: 2015-02-09, By Gavin
// Modified: 2015-02-09, By Gavin
// Link: https://tour.golang.org/flowcontrol/5

package main

import (
	"fmt"
	"math"
	"runtime"
)

// 函数：返回一个数的平方根
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// 函数：返回x的n次方
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// 主函数
func main() {
	// 调用平方根函数
	fmt.Println("2的平方根是：", sqrt(2), "-4的平方根是：", sqrt(-4))

	// 调用指数函数
	fmt.Println(
		pow(3, 2, 10), // 9
		pow(3, 3, 20), // 20
	)

	// 运行时，Swich 语句
	fmt.Print("Go runs on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
