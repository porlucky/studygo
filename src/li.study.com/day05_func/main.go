/*
 * @Author: chirs
 * @Date: 2021-05-07 15:29:46
 * @LastEditTime: 2021-05-08 08:44:21
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day05\main.go
 * @Description:func
 */

package main

import "fmt"

//常规函数
func f1(x int, y int) (ret int) {
	return x + y
}

//没有返回值
func f2(x int, y int) {
	fmt.Println(x + y)
}

//没有参数
func f3() int {
	return 3
}

//返回值未命名可以直接使用
func f4(x int, y int) (ret int) {
	ret = x + y
	return
}

//多个返回值
func f5() (int, string) {
	return 8858, "苹果哟"
}

//参数简写
func f6(x, y, z int, s, t string, r, f bool) int {
	return x + y
}

//可变长参数要写在最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) //y类型是q参数类型的切片
}
func main() {
	a := f1(1, 3)
	b := f3()
	c := f4(4, 5)
	fmt.Println(a)
	f2(21, 3)
	fmt.Println(b)
	fmt.Println(c)
	m, n := f5()
	fmt.Println(m)
	fmt.Println(n)
	_, x := f5()
	fmt.Println(x)
	ff := f6(1, 2, 3, "s", "f", true, false)
	fmt.Printf("x+y=%d", ff)
	f7("实时")

	e := [3]int{10, 20, 30}
	modifyArray(e) //在modify中修改的是a的副本x

	fmt.Println(e) //[10 20 30]

	qb := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(qb) //在modify中修改的是b的副本x
	fmt.Println(qb)  //[[1 1] [1 1] [1 1]]
}

func modifyArray(x [3]int) {
	x[0] = 100
	fmt.Println(x)
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
	fmt.Println(x)
}
