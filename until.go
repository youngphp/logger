package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

func CheckPathIsExits(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		fileName = file
		lineNo = line
		funcName = runtime.FuncForPC(pc).Name()
	}
	return
}

func GetLevelText(level int) (levelName string) {
	switch level {
	case DebugLevel:
		levelName = "DEBUG"
	case TraceLevel:
		levelName = "TRACE"
	case InfoLevel:
		levelName = "INFO"
	case WarnLevel:
		levelName = "WARN"
	case ErrorLevel:
		levelName = "ERROR"
	case FatalLevel:
		levelName = "FATAL"
	default:
		levelName = "UNKNOW"
	}
	return
}

func WriteLog(file *os.File, level int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)
	levelName := GetLevelText(level)
	fmt.Fprintf(file, "[%s][%s][%s---%s:%d] :%s\n", nowStr, levelName, fileName, funcName, lineNo, msg)
}
