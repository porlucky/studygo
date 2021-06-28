/*
 * @Author: chirs
 * @Date: 2021-05-15 08:48:41
 * @LastEditTime: 2021-05-17 15:54:17
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day17\main.go
 * @Description:转化JSON 序列化反序列化 结构体的复习
 */
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name string
	age  int
}
type addr struct {
	city string
	id   int
}

type student struct {
	name string
	addr
}

func main() {
	p1 := newperson("佩奇", 5)
	fmt.Println(p1)
	p1.look()
	p1.play("足球", 3)
	ret := p1.sum(5, 6)
	fmt.Println(ret)
	fmt.Println(p1.age)
	p1.addage()
	fmt.Println(p1.age)
	s1 := student{
		name: "小红",
		addr: addr{
			city: "大连",
			id:   23,
		},
	}
	s1.study()

	//JSON 序列化
	d := dog{"大黄", 15}
	ret1, err := json.Marshal(d)

	if err != nil {
		fmt.Println("出错误了！")
	} else {
		fmt.Println(string(ret1))
	}
	//JSON 反序列化
	str1 := `{"name":"猴子","age":5000}`
	var d2 dog
	err = json.Unmarshal([]byte(str1), &d2)
	if err != nil {
		fmt.Println("错误")
	} else {
		fmt.Println(d2)
	}

}

type dog struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//person的构造函数
func newperson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}
func (p *person) look() {
	fmt.Printf("%v在看书\n", p.name)
}

func (p *person) play(ball string, n int) {
	fmt.Printf("%v在玩%d个%v\n", p.name, n, ball)
}

//指针接收者
func (p *person) sum(m int, n int) int {
	return n + m
}
func (p *person) addage() {
	fmt.Printf("%T\n", p)
	p.age++
	fmt.Println(p.age)
}
func (s student) study() {
	fmt.Println(s.addr.city)
}
