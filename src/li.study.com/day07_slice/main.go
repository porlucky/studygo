/*
 * @Author: chirs
 * @Date: 2021-05-10 13:41:00
 * @LastEditTime: 2021-06-02 13:32:41
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day07_slice\main.go
 * @Description:切片
 */
package main

import "fmt"

func main() {

	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	f1(&a)
	// f2(a)
	// fmt.Println(a)
	// s := make([]int, 5, 19)
	// fmt.Printf("len:%v *** cap:%v\n", len(s), cap(s))
	a1 := [5]int{1, 2, 3, 4, 5}
	s := a1[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
}
func f1(x *[5]int) {
	x1 := *x
	fmt.Println(*x)
	x1[1] = 125
	fmt.Println(x1)
}
