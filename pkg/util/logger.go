package util

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var LogrusObj *logrus.Logger

func init() {
	src, err := setOutPutFile()
	if err != nil {
		panic(err)
	}
	if LogrusObj != nil {
		LogrusObj.Out = src
		return
	}
	logger := logrus.New()
	logger.Out = src                   //设置输出
	logger.SetLevel(logrus.DebugLevel) //日志级别
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}

func setOutPutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	// 获取当前路径
	if dir, err := os.Getwd(); err != nil {
		return nil, err
	} else {
		logFilePath = dir + "/logs/"
	}
	// 是否存在，不存在则创建
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	// format必须使用时间 `Mon Jan 2 15:04:05 MST 2006` 来指定给定时间/字符串的格式化/解析方式。
	logFileName := now.Format("2006-01-02") + ".log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		if _, err = os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return src, nil
}
