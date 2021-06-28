/*
 * @Author: chirs
 * @Date: 2021-05-10 16:33:01
 * @LastEditTime: 2021-05-11 13:12:13
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day08\main.go
 * @Description:
 */
package main

import "fmt"

func main() {
	m1 := make(map[int]string, 10)
	m1[1] = "张海山"
	m1[2] = "张翠山"

	if v, ok := m1[2]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("找不到")
	}

	if v1, ok := m1[4]; ok {
		fmt.Println(v1)
	} else {
		fmt.Println("找不到")
	}
}
