// Created: 2015-02-06, By Gavin
// Modified: 2016-12-15, By Gavin
// Modified: 2020-06-11, By Gavin
// Link: http://tour.golang.org/basics/4

// 包声明
package main

// 引入包
import (
	"fmt"
	"math"
)
// Go 语言数据类型
// https://mp.weixin.qq.com/s?__biz=MzI0NzQ0MDU0NA==&mid=2247483915&idx=1&sn=f286d112676307a12e25b24faeab9faa&scene=21#wechat_redirect

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
	// 变量定义
	// https://mp.weixin.qq.com/s?__biz=MzI0NzQ0MDU0NA==&mid=2247483918&idx=1&sn=f7435bddb9ef14f33503dc168f25dde5&scene=21#wechat_redirect
	// 定义整数
	var i_zhengshu int
	//定义浮点型
	var f_fudian float64
	//定义布尔型
	var b_buer bool
	//定义字符串
	var s_zifuchuan string
	//定义指针
	var a_zhizhen *int
	//定义数据
	var a_shuzu []int
	//格式化输出整形、浮点型、布尔型、字符串的默认值
	fmt.Printf("%v %v %v %q\n",i_zhengshu, f_fudian, b_buer, s_zifuchuan)
	//结果：0 0 false ""

	//非格式化输出去指针、数组的默认值
	fmt.Println(a_zhizhen, a_shuzu)
	//结果：<nil> []

	// 在 Go 程序中，一行代表一个语句结束。
	// 如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法。
	// Go 语言的字符串连接可以通过 + 实现
	fmt.Println("Hello, Gavin." + "Today is 2020-06-11.")
	fmt.Println(math.Pi)       // 输出Pi
	fmt.Println(addxy(42, 13)) // 调用加法函数
	fmt.Println(addxy(1234, 325)) // 调用加法函数

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
	var f float64 = math.Sqrt(float64(num1 * num1 + num2 * num2))
	var z int = int(f) // 强制类型转换
	fmt.Println(num1, "与", num2, "的平方根是：", z)
}
