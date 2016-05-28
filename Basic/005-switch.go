// Created: 2015-02-13, By Gavin
// Modified: 2015-02-13, By Gavin
// Link: https://tour.golang.org/flowcontrol/10

package main

import (
	"fmt"
	"time"
)

func main() {

	// Go语言的默认时间格式：2009-11-10 23:00:00 UTC
	t := time.Now()
	fmt.Println(t)
	// 2015-02-13 11:55:56.0869194 +0800 CST

	// 例1
	// Swich 语句：判断今天距离星期六还有几天
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// 例2
	// Switch语句，判断当前时间是早晨、中午或晚上
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening.")
	}

}
