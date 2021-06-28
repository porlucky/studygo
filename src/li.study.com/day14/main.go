/*
 * @Author: chirs
 * @Date: 2021-05-14 10:14:42
 * @LastEditTime: 2021-05-14 15:29:47
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day14\main.go
 * @Description:
 */
package main

import "fmt"

type student struct {
	name string
	age  int
}

func newStudent(name string, age int) *student {
	return &student{
		name: name,
		age:  age,
	}
}
func (s student) dream() {
	fmt.Printf("%v今年%d岁\n", s.name, s.age)
}
func (s student) de1name(name string) {
	s.name = name
}
func (s *student) dename(name string) {
	s.name = name
}
func main() {
	s1 := newStudent("二哈", 5)
	s1.de1name("小黄")
	fmt.Println(s1.name)
	s1.dename("大黄")
	fmt.Println(s1.name)
	s1.dream()
}
