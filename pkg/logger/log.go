// Package logger
// @Author: itcyy@HuaWei
// @File: infrastructure/pkg/logger/log.go
// @Time: 2023-12-01 22:02:43
package logger

import (
	"log"
	"os"
	"time"
)

var Logger *log.Logger

func CoonLog() {
	logName := "logs/" + time.Now().Format("2006-01-02") + "-" + "service.log"
	logFile, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err, "log error")
	}
	Logger = log.New(logFile, "INFO", log.Ldate|log.Ltime|log.Lshortfile|log.LstdFlags|log.Llongfile|log.Lmicroseconds|log.LUTC|log.Lmsgprefix|log.LUTC|log.LstdFlags)

}

// Info 详情
func Info(args ...interface{}) {
	CoonLog()
	Logger.SetPrefix("[Qxz-Server] [INFO]")
	Logger.Println(args...)
	log.Println(args...)
}

// Danger 错误 为什么不命名为 error？避免和 error 类型重名
func Danger(args ...interface{}) {
	CoonLog()
	Logger.SetPrefix("[Qxz-Server] [ERROR]")
	log.Println(args...)
	Logger.Fatal(args...)

}

// Warning 警告
func Warning(args ...interface{}) {
	CoonLog()
	Logger.SetPrefix("[Qxz-Server] [WARNING]")
	Logger.Println(args...)
	log.Println(args...)
}

// DeBug debug
func DeBug(args ...interface{}) {
	CoonLog()
	Logger.SetPrefix("[Qxz-Server] [DeBug]")
	log.Println(args...)
	Logger.Println(args...)
}
