// Created: 2020-06-22, By Gavin
// Created: 2020-06-22, By Gavin
// Link: https://www.cnblogs.com/tonylp/archive/2013/03/20/2971272.html

// 包声明
package main

// 导入包
import (
	"fmt"
)

// 主函数
func main() {
	// Golang 斐波那契数列
	fmt.Println("输入两个正整数m和n，求其最大公约数和最小公倍数。")
	// 输入要计算的项数j（输入10，计算到第9项）
	fmt.Println("请输入第1个整数:")
	var num1 int
	fmt.Scanln(&num1)

	fmt.Println("请输入第2个整数:")
	var num2 int
	fmt.Scanln(&num2)
	fmt.Println("您输入的数字为：",num1, num2)
	var first, second = num1, num2
	var temp int
	// 如果num1小于num2，则进行交换
	if num1 < num2 {
		temp = num1
		num1 = num2
		num2 = temp
	}
	// 辗转相除法
	for num2 != 0 {
		temp = num1 % num2
		num1 = num2
		num2 = temp
	}
	// 输出最大公约数
	fmt.Println("两个数的最大公约数为：", num1)
	// 输出最小公倍数
	fmt.Println("两个数的最小公倍数为：", first*second/num1)
}
