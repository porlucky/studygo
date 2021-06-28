/*
 * @Author: chirs
 * @Date: 2021-05-17 16:11:33
 * @LastEditTime: 2021-05-18 10:15:44
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day18\do.go
 * @Description:
 */
package main

import (
	"fmt"
)

type students struct {
	id   int64
	name string
}
type stumar struct {
	allstudent map[int64]students
}

//添加
func (s stumar) addone() {
	var no int64
	var name string
	fmt.Print("请输入学号")
	fmt.Scanln(&no)
	_, ok := onemaster.allstudent[no]
	if ok {
		fmt.Println("学号存在，重新输入！")
		fmt.Scanln(&no)
	}
	fmt.Print("请输入姓名")
	fmt.Scanln(&name)
	s.allstudent[no] = newStudents(no, name)
	fmt.Println("添加成功")
}

//查询
func (s stumar) selectone() {
	for _, v := range s.allstudent {
		fmt.Printf("学号：%d 姓名：%s\n", v.id, v.name)
	}
}

//修改
func (s stumar) modifyone() {
	var no int64
	fmt.Print("请输入学号")
	fmt.Scanln(&no)
	objv, ok := s.allstudent[no]
	if ok {
		fmt.Println(objv)
		var nname string
		fmt.Print("请输入新名字：")
		fmt.Scanln(&nname)
		objv.name = nname
		s.allstudent[no] = objv
	} else {
		fmt.Println("学号不存在")
		return
	}
}

//删除
func (s stumar) delone() {
	var no int64
	fmt.Print("请输入要删除的学号")
	fmt.Scanln(&no)
	_, ok := s.allstudent[no]
	if !ok {
		fmt.Println("学号不存在，重新输入！")
		return
	}
	delete(s.allstudent, no)
}

func newStudents(id int64, name string) students {
	return students{
		id:   id,
		name: name,
	}
}
