/*
 * @Author: chirs
 * @Date: 2021-05-27 10:04:53
 * @LastEditTime: 2021-05-27 10:50:54
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\day20\calc\calc.go
 * @Description:
 */
package calc

//包中的标识符（变量名、函数名、接口、结构体）如果首字母是小写的表示私有（只能自己包内调用）
//外部调用必须首字母大写
func Add(x int, y int) int {
	return x + y
}
