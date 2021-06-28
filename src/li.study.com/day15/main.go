/*
 * @Author: chirs
 * @Date: 2021-05-14 16:38:05
 * @LastEditTime: 2021-05-15 08:33:11
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day15\main.go
 * @Description:
 */
package main

import "fmt"

type hua struct {
	color string
}

func (a hua) kai() {
	fmt.Printf("%v花生长\n", a.color)
}

type meigui struct {
	add    string
	isopen int
	hua
}

func (m meigui) who() {
	//m.kai()
	fmt.Printf("生长在%v开了\n", m.add)
}
func main() {
	m := meigui{
		add:    "红玫瑰",
		isopen: 1,
		hua: hua{
			color: "红",
		},
	}
	m.who()
	m.kai()
}
