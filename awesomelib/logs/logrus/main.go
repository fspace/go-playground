package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	// 以JSON格式为输出，代替默认的ASCII格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// 以Stdout为输出，代替默认的stderr
	logrus.SetOutput(os.Stdout)
	// 设置日志等级
	logrus.SetLevel(logrus.WarnLevel)
}
func main() {
	//logrus.WithFields(logrus.Fields{
	//	"animal": "walrus",
	//	"size":   10,
	//}).Info("A group of walrus emerges from the ocean")
	//
	//logrus.WithFields(logrus.Fields{
	//	"omg":    true,
	//	"number": 122,
	//}).Warn("The group's number increased tremendously!")
	//
	//logrus.WithFields(logrus.Fields{
	//	"omg":    true,
	//	"number": 100,
	//}).Fatal("The ice breaks!")

	// -----------
	demo2()
}

func demo2() {

	var log = logrus.New()

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
	file.Close()
}
