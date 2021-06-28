/*
 * @Author: chirs
 * @Date: 2021-05-25 10:05:35
 * @LastEditTime: 2021-05-25 10:52:25
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day19_interface\interface05\main.go
 * @Description:接口的值接受者和指针接收者 值接收者既能接受指针也能接受值 指针接收者不能接受值只能接受指针
 */
package main

import "fmt"

type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int8
}

//值接受者实现的方法
// func (c cat) move() {
// 	fmt.Println("猫步")
// }

// func (c cat) eat(food string) {
// 	fmt.Printf("%v猫吃%v...\n", c.name, food)
// }

//值接受者实现的方法
func (c *cat) move() {
	fmt.Println("猫步")
}

func (c *cat) eat(food string) {
	fmt.Printf("%v猫吃%v...\n", c.name, food)
}
func main() {
	var a1 animal
	c1 := cat{"tom", 4}
	fmt.Println(c1)
	c2 := &cat{"mary", 4}
	fmt.Println(c2)
	a1 = &c1
	a1 = c2
	a1.eat("鱼")
}
