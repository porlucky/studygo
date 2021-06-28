/*
 * @Author: chirs
 * @Date: 2021-05-13 08:45:02
 * @LastEditTime: 2021-05-14 09:50:16
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day13\main.go
 * @Description:
 */
package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}
type book struct {
	page   int
	name   string
	author string
}

//book构造函数
func newbook(page int, name string, author string) book {
	return book{
		page:   page,
		name:   name,
		author: author,
	}
}

func (b book) look() {
	fmt.Printf("《%v》精装版一共%v页，作者：%v", b.name, b.page, b.author)
}

//函数传值永远是copy
func f(p *person) {
	(*p).gender = "女" //语法糖可以像下面那么写，go不能操作指针，系统会自动判断
	p.gender = "无知"
}

func main() {
	var p person
	p.name = "优势"
	p.age = 16
	p.gender = "男"
	p.hobby = []string{"跳水", "跑步"}
	f(&p) //传递内存地址
	var b = newbook(100, "减少", "大迷糊")
	b.look()

}
