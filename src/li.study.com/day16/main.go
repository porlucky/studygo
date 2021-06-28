/*
 * @Author: chirs
 * @Date: 2021-05-15 08:48:41
 * @LastEditTime: 2021-05-15 13:14:54
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day16\main.go
 * @Description:转化JSON 格式
 */

package main

import (
	"encoding/json"
	"fmt"
)

//人的结构体
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "一个人",
		Age:  1,
	}
	//序列化
	x, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("Marshal格式错误")
	}
	fmt.Printf("%#v", string(x))
	str := `{"name":"张飞","age":15}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Println(p2)
}
