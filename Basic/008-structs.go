// Created: 2015-02-15, By Gavin
// Modified: 2015-03-03, By Gavin
// Link: https://tour.golang.org/moretypes/2

package main

import "fmt"

// 定义结构体
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2} // v1类型为 Vertex
	v2 = Vertex{X: 1} // 默认情况下 Y:0
	v3 = Vertex{}     // p类型为 *Vertex
)

func main() {
	// 结构体调用
	// fmt.Println(Vertex{1, 2})

	fmt.Println(v1)

	v1.X = 4
	fmt.Println(v1.X)

	// 指向结构体的指针
	p := &v1
	p.X = 1e9
	fmt.Println(v1)

	// 结构体常量：Struct Literals
	fmt.Println(v1, p, v2, v3)

}
