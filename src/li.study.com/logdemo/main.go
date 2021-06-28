/*
 * @Author: chirs
 * @Date: 2021-06-02 14:09:46
 * @LastEditTime: 2021-06-03 12:54:07
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\logdemo\main.go
 * @Description:
 */
/*
*日志级别
	1.Debug
	2.Trace
	3.Info
	4.Warning
	5.Error
	6.Fatal
*/

package main

import (
	"time"

	"li.study.com/logout"
)

func main() {
	logouter := logout.NewLog("info")
	for {
		logouter.Debug("这是一条Debug")
		logouter.Info("这是一条Info")
		logouter.Worning("这是一条Worning")
		logouter.Err("这是一条Err")
		time.Sleep(3 * time.Second)
	}

}
