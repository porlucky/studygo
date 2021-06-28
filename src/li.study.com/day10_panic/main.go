/*
 * @Author: chirs
 * @Date: 2021-05-11 16:00:20
 * @LastEditTime: 2021-05-12 10:02:28
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day10\main.go
 * @Description:
 */
package main

import (
	"fmt"
)

func funcA() {
	fmt.Println("func A")
}

func funcB() {
	defer func() {
		err := recover()
		fmt.Printf("err:%v\n", err)
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func funcC() {
	fmt.Println("func C")
}
func main() {
	funcA()
	funcB()
	funcC()
}
