/*
 * @Author: chirs
 * @Date: 2021-05-27 13:44:21
 * @LastEditTime: 2021-05-27 15:18:41
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day20\openfile\main.go
 * @Description:读文件
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func openfile() {
	fileobj, err := os.Open("../test.txt")
	if err != nil {
		fmt.Printf("open failed err：%v", err)
		return
	}
	//延迟关闭
	defer fileobj.Close()
	var tmp = make([]byte, 128)
	for {
		n, err := fileobj.Read(tmp)
		if err != nil {
			fmt.Printf("read failed err%v", err)
			return
		}
		fmt.Print(string(tmp))
		if n < 128 {
			return
		}
	}
}
func bufiofild() {
	fileobj, err := os.Open("../test.txt")
	if err != nil {
		fmt.Printf("open failed err：%v", err)
		return
	}
	//延迟关闭
	defer fileobj.Close()
	reader := bufio.NewReader(fileobj)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("open bufio err：%v", err)
			return
		}
		fmt.Print(str)
	}

}
func readuiufio() {
	str, err := ioutil.ReadFile("../test.txt")
	if err != nil {
		fmt.Printf("field err：%v", err)
		return
	}
	fmt.Print(string(str))
}
func main() {
	//openfile()
	//bufiofild()
	readuiufio()
}
