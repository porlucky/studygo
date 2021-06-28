/*
 * @Author: chirs
 * @Date: 2021-05-18 11:01:42
 * @LastEditTime: 2021-05-18 13:39:50
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day19_interface\interface02\main.go
 * @Description:
 */
package main

import "fmt"

type car interface {
	run()
}
type che1 struct {
	brand string
}
type che2 struct {
	brand string
}

func (c1 che1) run() {
	fmt.Printf("%v跑的快\n", c1.brand)
}
func (c2 che2) run() {
	fmt.Printf("%v跑的快\n", c2.brand)
}

//接受car 类型的变量，，只要有run()方法都能传递
func drive(c3 car) {
	c3.run()
}
func main() {
	var f1 = che1{
		brand: "法拉利",
	}
	var f2 = che2{
		brand: "保斯捷",
	}
	drive(f1)
	drive(f2)
}
