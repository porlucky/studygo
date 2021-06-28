/*
 * @Author: chirs
 * @Date: 2021-05-28 16:21:27
 * @LastEditTime: 2021-05-31 13:49:18
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day21_time\main.go
 * @Description:time包
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// //fmt.Println(now)
	// fmt.Println(now.Year())
	// fmt.Println(now.Month())
	// fmt.Println(now.Day())
	// fmt.Println(now.Hour())
	// fmt.Println(now.Minute())
	// fmt.Println(now.Second())
	// //时间戳
	// fmt.Println(now.Unix())
	// //+一小时
	// fmt.Println(now.Add(1 * time.Hour))
	// //时间格式化
	// /*
	// 	y     m      d	    h    m    s
	// 	2006  1      2      3    4    5
	// */
	// fmt.Println(now.Format("2006-01-02"))
	// fmt.Println(now.Format("2006/01/02 15:04:05"))
	// timeobj, err := time.Parse("2006-01-02 15:04:05", "2021-01-20 12:24:21")
	// if err != nil {
	// 	fmt.Println("传入错误")
	// 	return
	// }
	// fmt.Println(timeobj)
	// fmt.Println(timeobj.Unix())
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2021/06/01 13:49:00", loc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("******************\n")
	fmt.Println(timeObj.Sub(now))
	fmt.Println(now)
	//fmt.Println(timeObj.Sub(now))

}
