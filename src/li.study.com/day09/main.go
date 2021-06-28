/*
 * @Author: chirs
 * @Date: 2021-05-11 13:53:42
 * @LastEditTime: 2021-05-11 16:00:00
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day09\main.go
 * @Description:
 */
package main

import "fmt"

//定义了一个calc函数 函数有三个参数分别是 x int y int  func op  op函数有两个参数一个返回值 calc函数返回一个int值
func calc(x, y int, op func(int, int) int) int {
	ret := op(x, y)
	fmt.Println(ret)
	return ret
}
func add(x, y int) int {
	return x + y
}

func main() {
	ret2 := calc(10, 20, add)
	fmt.Println(ret2)
	c := add                        // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c

	f := add                        // 将函数add赋值给变量f1
	fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	fmt.Println(f(10, 20))          // 像调用add一样调用f
}
