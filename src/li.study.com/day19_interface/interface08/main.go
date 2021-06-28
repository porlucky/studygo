/*
 * @Author: chirs
 * @Date: 2021-05-27 09:25:41
 * @LastEditTime: 2021-05-27 09:34:31
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day19_interface\interface08\main.go
 * @Description:
 */
package main

import "fmt"

func assert(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("字符串:%v", v)
	case int:
		fmt.Printf("int:%v", v)
	case int64:
		fmt.Printf("int64:%v", v)
	case bool:
		fmt.Printf("boll%v", v)
	}

}
func main() {
	//assert("oksss")
	//assert(5)
	assert(int64(5))
}
