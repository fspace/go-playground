package hilogrus

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到指定文件（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logFile := "app.log"
	file, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(file)

	// 设置只记录日志级别为warn及其以上的日志
	log.SetLevel(log.WarnLevel)
}

func Main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}
