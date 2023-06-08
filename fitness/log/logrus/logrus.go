package logrus

import (
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func InitLog() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(true)

	// 检查日志目录是否存在，不存在则创建
	logDir := filepath.Join("..", "fitness", "log", "logrus", "logFile")
	if err := os.MkdirAll(logDir, 0644); err != nil {
		logrus.Fatalf("Failed to create log directory: %v", err)
	}

	// 设置日志输出到文件
	logFilename := filepath.Join(logDir, "log.txt")
	f, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(f)
	} else {
		logrus.Infof("Failed to log to file, using default stderr: %v", err)
	}

	// 设置日志级别为Debug以上
	logrus.SetLevel(logrus.DebugLevel)
}
