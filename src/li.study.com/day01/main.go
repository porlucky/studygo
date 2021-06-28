/*
 * @Author: chirs
 * @Date: 2021-04-24 15:21:37
 * @LastEditTime: 2021-04-26 09:19:35
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day01\main.go
 * @Description:
 */
package main

import "fmt"

var (
	name string
	age  int
)

func main() {
	name = "阿飛"
	age = 15
	s1 := "三十分" //在函数里使用 不能在外面使用
	fmt.Println(age)
	fmt.Printf("name:%s\n", name)
	fmt.Println(s1)
}
