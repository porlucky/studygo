/*
 * @Author: chirs
 * @Date: 2021-05-18 10:17:55
 * @LastEditTime: 2021-05-18 13:57:43
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day19_interface\interface01\main.go
 * @Description:接口是一种类型，规定了变量有哪些方法 不关心变量类型，只关心调用的方法存在不存在
 */
//接口规定了
package main

import "fmt"

type speaker interface {
	speak() //只要有这个方法的都可以传，方法签名
}
type cat struct{}
type dog struct{}

func (c cat) speak() {
	fmt.Println("miaomiaomiao")
}
func (d dog) speak() {
	fmt.Println("wangwangwang")
}
func da(s speaker) {
	s.speak()
}
func main() {
	var c1 cat
	var d1 dog
	da(c1)
	da(d1)
}
