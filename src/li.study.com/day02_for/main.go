/*
 * @Author: chirs
 * @Date: 2021-04-28 09:58:08
 * @LastEditTime: 2021-05-18 08:45:17
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day02\main.go
 * @Description:
 */
package main

import "fmt"

func main() {
	var a1 = [3]string{"hh", "aa", "bb"}
	fmt.Println(a1)
	for _, v := range a1 {
		println(v)
	}

	var b1 = [...][2]int{{1, 2}, {1, 3}, {1, 4}, {5, 6}}
	for _, v := range b1 {
		fmt.Println(v)
		for _, w := range v {
			fmt.Println(w)
		}
	}

	c1 := [...]int{1, 3, 5, 7, 9}
	for i := 0; i < len(c1); i++ {
		for j := i + 1; j < len(c1); j++ {
			if c1[i]+c1[j] == 8 {
				fmt.Println(i, j)
			}
		}
	}
}
