/*
 * @Author: chirs
 * @Date: 2021-05-06 15:48:10
 * @LastEditTime: 2021-05-07 09:20:20
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day03\main.go
 * @Description:
 */
package main

import "fmt"

func main() {
	s3 := []int{1, 3, 5}
	s3[2] = 2000
	fmt.Println(s3)
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	for _, v := range s3 {
		fmt.Println(v)
	}
}
