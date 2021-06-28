/*
 * @Author: chirs
 * @Date: 2021-05-07 09:20:17
 * @LastEditTime: 2021-05-07 15:27:28
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day04\main.go
 * @Description:append
 */
package main

import "fmt"

func main() {
	//map 的切片
	var s1 = make([]map[int]string, 5, 10)
	s1[0] = make(map[int]string) //初始化map
	s1[0][100] = "a"
	fmt.Println(s1)

	//切片的map
	var m1 = make(map[int][]string)
	m1[0] = []string{"hhh", "aaaa"}
	fmt.Println(m1)
}
