/*
 * @Author: chirs
 * @Date: 2021-05-12 10:28:56
 * @LastEditTime: 2021-05-12 15:53:48
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day11\main.go
 * @Description:
 */
package main

import "fmt"

/* 分金币 */
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

//循环map 返回剩下的金币和map数组
func dispatchCoin(x int) (n int, m map[string]int) {
	total := 0
	for _, name := range users {
		distribution[name] = ts(name)
		total += ts(name)
	}
	ret := x - total
	return ret, distribution
}

//判断名字有几个对应的字母，返回名字得到的金币
func ts(s string) int {
	a := 0
	for _, v := range s {
		if v == 'e' || v == 'E' {
			a += 1
		}
		if v == 'i' || v == 'I' {
			a += 2
		}
		if v == 'o' || v == 'O' {
			a += 3
		}
		if v == 'u' || v == 'U' {
			a += 4
		}
	}
	return a
}

func main() {
	left := 0
	var m = make(map[string]int, len(users))
	left, m = dispatchCoin(coins)
	fmt.Println("剩下：", left)
	fmt.Println("数组", m)
}
