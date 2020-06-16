// Created: 2020-06-16, By Gavin
// Created: 2020-06-16, By Gavin
// Link: https://blog.csdn.net/zhangyexinaisurui/article/details/81174561

// 包声明
package main

// 导入包
import(
	"fmt"
)

// 递归函数
func rescuive(i int) int{
	if i<=1{
		return 1
	}
	return rescuive(i-1)+rescuive(i-2)
}

// 主函数
func main() {
	// Golang 斐波那契数列
	fmt.Println("斐波那契数列，请输入一个小于30的整数：")
	// 输入要计算的项数j（输入10，计算到第9项）
	var j int
	fmt.Scanln(&j)
	// 循环
	for i:=1;i<j;i++{
		// 调用递归函数
		sum:=rescuive(i)
		// 输出结果
		// fmt.Println(sum)
		fmt.Println(i,sum)
	}
}
