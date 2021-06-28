/*
 * @Author: chirs
 * @Date: 2021-06-02 14:17:26
 * @LastEditTime: 2021-06-07 10:15:35
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\logout\file.go
 * @Description:
 */
/*日志级别
1.Debug
2.Trace
3.Info
4.Warning
5.Error
6.Fatal
*/
package logout

//结构体
type Txter struct {
	level       logLevel
	filepath    string //路径
	filename    string //文件名
	maxfilesize int64
}

//构造函数
func NewTxter(level, fp, fn string, maxfilesize int64) *Txter {
	return &Txter{}
}
func (t Txter) Debug(msg string) {

}
func (t Txter) Info(msg string) {

}
func (t Txter) Worning(msg string) {

}
func (t Txter) Err(msg string) {

}
