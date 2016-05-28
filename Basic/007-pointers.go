// Created: 2015-02-15, By Gavin
// Modified: 2015-02-15, By Gavin
// Link: https://tour.golang.org/moretypes/1

package main

import "fmt"

func main() {
	i, j := 42, 2701

	// 指针保存的是变量的内存地址
	p := &i // 定义指针变量，指向地址i
	fmt.Println("The address of pointer i is: ", p)
	fmt.Println("The value of i is: ", *p)              // 指向i的值
	*p = 21                                             // 通过指针为i赋值
	fmt.Println("After modify, the value of i is: ", i) // 显示i的新值

	p = &j                                              // 指向变量j
	*p = *p / 37                                        // 通过指针对j做除法
	fmt.Println("After devide, the value of j is: ", j) // 显示j的值

}
