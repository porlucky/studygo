/*
 * @Author: chirs
 * @Date: 2021-05-12 15:54:07
 * @LastEditTime: 2021-05-18 13:29:52
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day12\main.go
 * @Description:结构体
 */
package main

import "fmt"

type myint int

type person struct {
	name string
	age  int
}

func main() {
	var n myint = 100
	fmt.Println(n)
	fmt.Printf("%T", n)
	var p1 = person{
		name: "hh",
		age:  15,
	}
	fmt.Println(p1)
}
