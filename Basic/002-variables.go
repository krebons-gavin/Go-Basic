// Created: 2015-02-07, By Gavin
// Modified: 2015-02-09, By Gavin
// Link: http://tour.golang.org/basics/9

package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 数字常量：数字常量是高精密度的
const (
	Big   = 1 << 100  // 1左移100位，相当于乘以2的100次方
	Small = Big >> 99 // Big右移99位，相当于Small = 2
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// 定义整型变量
	// 注意：go 是静态类型的语言
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"

	// 不用指定变量类型的定义，系统会根据初始赋值自己确定类型
	k := 3 // 等价于 var k int = 3

	fmt.Println(i, j, k, c, python, java)
	// 结果：1 2 3 true false no!

	// 注意：Println不支持，Printf才支持%式的输出
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))

	fmt.Printf("%t\n", 1 == 2)
	fmt.Printf("255的二进制：%b\n", 255)
	fmt.Printf("255的八进制：%o\n", 255)
	fmt.Printf("浮点数-Pai：%f\n", math.Pi)

	// 定义数组
	var a [5]int
	fmt.Println("array a:", a)
	a[1] = 10
	a[3] = 30
	fmt.Println("assign:", a)
	fmt.Println("len:", len(a))

	// 常量：可以是字符型, string, boolean, 或者数字
	// 常量不能用 := 语法来声明
	const World = "世界"
	fmt.Println("这是const常量的例子：hello", World)

	// 定义复数
	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	// 数字常量
	fmt.Println("数字常量例子：")
	fmt.Println(needInt(Small)) // 21
	// fmt.Println(needInt(Big))
	// constant 1267650600228229401496703205376 overflows int
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big))   // 1.2676506002282295e+29

}
