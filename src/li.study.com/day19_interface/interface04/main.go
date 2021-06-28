/*
 * @Author: chirs
 * @Date: 2021-05-18 15:59:46
 * @LastEditTime: 2021-05-18 16:19:29
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day19_interface\interface04\main.go
 * @Description:
 */
package main

import "fmt"

type People interface {
	Speak(string) string
}
type Mover interface {
	move()
}

type dog struct{}
type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func (d *dog) move() {
	fmt.Println("狗会动")
}

func main() {
	var x Mover
	var fugui = &dog{} // 富贵是*dog类型
	x = fugui          // x可以接收*dog类型
	x.move()
	var peo People = &Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
