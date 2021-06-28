/*
 * @Author: chirs
 * @Date: 2021-06-02 14:16:20
 * @LastEditTime: 2021-06-03 14:48:30
 * @LastEditors: chirs
 * @FilePath: \companyd:\workspace\src\li.study.com\logout\console.go
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

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//Consoler
type logLevel uint16

const (
	UNKNOW logLevel = iota
	DEBUG
	Info
	Worning
	Err
)

//结构体
type Consoler struct {
	Level logLevel
}

//结构体构造函数
func NewLog(Levelstr string) Consoler {
	level, err := paseLoglevel(Levelstr)
	if err != nil {
		panic(err)
	}
	return Consoler{
		Level: level,
	}
}

func paseLoglevel(s string) (logLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "info":
		return Info, nil
	case "worning":
		return Worning, nil
	case "err":
		return Err, nil
	default:
		err := errors.New("无效级别")
		return UNKNOW, err
	}
}

func (c Consoler) Debug(msg string) {
	if c.Level <= DEBUG {
		now := time.Now()
		now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s][Debug] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}

}
func (c Consoler) Info(msg string) {
	if c.Level <= Info {
		now := time.Now()
		now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s][Info] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (c Consoler) Worning(msg string) {
	if c.Level <= Worning {
		now := time.Now()
		now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s][Worning] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
func (c Consoler) Err(msg string) {
	if c.Level <= Err {
		now := time.Now()
		now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s][Err] %s\n", now.Format("2006-01-02 15:04:05"), msg)
	}
}
