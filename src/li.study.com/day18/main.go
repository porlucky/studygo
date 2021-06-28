/*
 * @Author: chirs
 * @Date: 2021-05-17 15:57:12
 * @LastEditTime: 2021-05-18 09:01:30
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day18\main.go
 * @Description:学习系统
 */
package main

import (
	"fmt"
	"os"
)

var onemaster stumar

func showMenu() {
	fmt.Println("welcome!")
	fmt.Println(`
	1.添加人员
	2.查询所有
	3.修改人员
	4.删除人员
	5.退出
	`)
}
func main() {
	onemaster = stumar{
		allstudent: make(map[int64]students),
	}
	for {
		showMenu()
		fmt.Print("请输入!")
		var n int
		fmt.Scanln(&n)
		if n == 1 {
			onemaster.addone()
		} else if n == 2 {
			onemaster.selectone()
		} else if n == 3 {
			onemaster.modifyone()
		} else if n == 4 {
			onemaster.delone()
		} else if n == 5 {
			os.Exit(1)
		} else {
			fmt.Print("错误")
		}
	}
}
