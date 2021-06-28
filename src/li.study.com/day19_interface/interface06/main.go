/*
 * @Author: chirs
 * @Date: 2021-05-25 10:23:33
 * @LastEditTime: 2021-05-25 10:50:35
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day19_interface\interface06\main.go
 * @Description:接口嵌套
 */
package main

import "fmt"

//接口嵌套。差不多就是类的继承 animal 嵌套 mover eater 两个接口拥有两个方法move eat
type animal interface {
	mover
	eater
}
type mover interface {
	move()
}
type eater interface {
	eat(string)
}

type cat struct {
	name string
	foot int8
}

func (c *cat) move() {
	fmt.Println("跑了")
}
func (c *cat) eat(food string) {
	fmt.Printf("%v吃le %v", c.name, food)
}
func main() {
	var a animal
	c1 := cat{"tom", 4}
	a = &c1
	a.eat("鱼")
	a.move()
}
