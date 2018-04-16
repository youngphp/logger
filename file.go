package logger

import (
	"fmt"
	"os"
	"time"
)

type FileLogger struct {
	level       int
	logPath     string
	file        *os.File
	warnFile    *os.File
	chanLogData chan *LogData
}

func NewFileLogger(level int, logPath string) LoggerInterface {
	logger := &FileLogger{
		level:       level,
		logPath:     logPath,
		chanLogData: make(chan *LogData, 100000),
	}
	logger.init()
	return logger
}

func (f *FileLogger) init() {
	//创建日志文件路径
	flag, err := CheckPathIsExits(f.logPath)

	if err != nil {
		panic(fmt.Sprintf("logpath is wrong:%s, the err is %v", f.logPath, err))
	}
	if !flag {
		err = os.MkdirAll(f.logPath, 0755)
		if err != nil {
			panic(fmt.Sprintf("create dir failed the err is %v", err))
		}
	}
	nowStr := fmt.Sprintf("%s", time.Now().Format("2006-01-02"))
	filename := fmt.Sprintf("%s/%s.log", f.logPath, nowStr)
	file, fileerr := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if fileerr != nil {
		panic(fmt.Sprintf("open file is failed:%s err is %v", filename, fileerr))
	}
	f.file = file
	filename = fmt.Sprintf("%s/%s-warn.log", f.logPath, nowStr)
	file, fileerr = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if fileerr != nil {
		panic(fmt.Sprintf("open file is failed:%s err is %v", filename, fileerr))
	}
	f.warnFile = file
	go f.WriteBackGroudLog()
}

func (f *FileLogger) SetLevel(level int) {
	if level < DebugLevel || level > FatalLevel {
		f.level = DebugLevel
	}
	f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > DebugLevel {
		return
	}
	logdata := WriteLog(DebugLevel, format, args...)
	select {
	case f.chanLogData <- logdata:
	default:
	}
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > TraceLevel {
		return
	}
	logdata := WriteLog(TraceLevel, format, args...)
	select {
	case f.chanLogData <- logdata:
	default:
	}
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > InfoLevel {
		return
	}
	logdata := WriteLog(InfoLevel, format, args...)
	select {
	case f.chanLogData <- logdata:
	default:
	}
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > WarnLevel {
		return
	}
	logdata := WriteLog(WarnLevel, format, args...)
	select {
	case f.chanLogData <- logdata:
	default:
	}
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > ErrorLevel {
		return
	}
	logdata := WriteLog(ErrorLevel, format, args...)
	select {
	case f.chanLogData <- logdata:
	default:
	}
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > FatalLevel {
		return
	}
	logdata := WriteLog(FatalLevel, format, args...)
	select {
	case f.chanLogData <- logdata:
	default:
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}

func (f *FileLogger) WriteBackGroudLog() {
	for logdata := range f.chanLogData {
		file := f.file
		if logdata.WarnAndFatal {
			file = f.warnFile
		}
		fmt.Fprintf(file, "[%s][%s][%s---%s:%d] :%s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
	}
}
