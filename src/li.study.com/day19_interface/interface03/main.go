/*
 * @Author: chirs
 * @Date: 2021-05-18 13:58:45
 * @LastEditTime: 2021-05-18 14:48:53
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day19_interface\interface03\main.go
 * @Description:
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
type chicken struct {
	feet int8
}

func (c1 cat) move() {
	fmt.Println("我会跑")
}
func (c1 cat) eat(s string) {
	fmt.Printf("%s在吃%v\n", c1.name, s)
}

func (c2 chicken) move() {
	fmt.Println("我会跑")
}
func (c2 chicken) eat(s string) {
	fmt.Printf("在吃%v", s)
}
func main() {
	var a1 animal
	bc := cat{
		name: "蓝猫",
		feet: 4,
	}
	a1 = bc
	a1.eat("鱼")
	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	a1.move()
	a1.eat("汉堡")
}
