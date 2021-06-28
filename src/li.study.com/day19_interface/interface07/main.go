/*
 * @Author: chirs
 * @Date: 2021-05-25 10:52:50
 * @LastEditTime: 2021-05-27 09:08:13
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day19_interface\interface07\main.go
 * @Description:空接口
 */
package main

import "fmt"

func show(a interface{}) {
	fmt.Printf("type:%T value:%v", a, a)

}
func main() {
	m1 := make(map[string]interface{}, 16)
	m1["name"] = "王丹"
	m1["age"] = 16
	m1["hobby"] = [...]string{"呵呵", "嘿嘿", "哈哈"}
	show(m1)
}
