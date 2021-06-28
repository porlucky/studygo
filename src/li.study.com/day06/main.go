/*
 * @Author: chirs
 * @Date: 2021-05-08 08:40:01
 * @LastEditTime: 2021-05-10 14:10:46
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day06\main.go
 * @Description:
 */
package main

import "fmt"

func main() {
	a := [...]string{"4-10", "4-09", "4-08", "4-07"}
	var b [6]string
	for i := 0; i < len(a); i++ {
		b[i] = a[len(a)-i-1]
	}
	fmt.Println(b)
}
