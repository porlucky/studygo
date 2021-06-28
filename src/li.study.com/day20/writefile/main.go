/*
 * @Author: chirs
 * @Date: 2021-05-27 15:20:25
 * @LastEditTime: 2021-05-28 09:51:05
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day20\writefile\main.go
 * @Description:写文件
 */
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func demo1() {
	fileobj, err := os.OpenFile("../test1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("ero:%v", err)
		return
	}
	defer fileobj.Close()
	n, err := fileobj.WriteString("这是我写的")
	if err != nil {
		fmt.Printf("ero:%v", err)
	}
	fmt.Println(n)
}
func demo2() {
	fileobj, err := os.OpenFile("../test.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("ero:%v", err)
		return
	}
	defer fileobj.Close()
	wr := bufio.NewWriter(fileobj)
	wr.WriteString("bufio 写入文件是写入缓存中的！")
	wr.Flush()
}
func demo3() {
	err := ioutil.WriteFile("../test.txt", []byte("ioutil写入文件"), 7777)
	if err != nil {
		fmt.Printf("ero:%v", err)
		return
	}
}
func main() {
	//demo1()
	//demo2()
	demo3()
}
