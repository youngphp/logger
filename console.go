package logger

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	level int
}

// 0:debug 1:trace 2:info 3:warn 4:error 5:fatal
func NewConsoleLogger(level int) LoggerInterface {
	logger := &ConsoleLogger{
		level: level,
	}
	return logger
}

func (c *ConsoleLogger) SetLevel(level int) {
	if level < DebugLevel || level > FatalLevel {
		c.level = DebugLevel
	}
	c.level = level
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	if c.level > DebugLevel {
		return
	}
	logdata := WriteLog(DebugLevel, format, args...)
	fmt.Fprintf(os.Stdout, "[%s][%s][%s---%s:%d] :%s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > TraceLevel {
		return
	}
	logdata := WriteLog(TraceLevel, format, args...)
	fmt.Fprintf(os.Stdout, "[%s][%s][%s---%s:%d] :%s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > InfoLevel {
		return
	}
	logdata := WriteLog(InfoLevel, format, args...)
	fmt.Fprintf(os.Stdout, "[%s][%s][%s---%s:%d] :%s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > WarnLevel {
		return
	}
	logdata := WriteLog(WarnLevel, format, args...)
	fmt.Fprintf(os.Stdout, "[%s][%s][%s---%s:%d] :%s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > ErrorLevel {
		return
	}
	logdata := WriteLog(ErrorLevel, format, args...)
	fmt.Fprintf(os.Stdout, "[%s][%s][%s---%s:%d] :%s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > FatalLevel {
		return
	}
	logdata := WriteLog(FatalLevel, format, args...)
	fmt.Fprintf(os.Stdout, "[%s][%s][%s---%s:%d] :%s\n", logdata.TimeStr, logdata.LevelStr, logdata.FileName, logdata.FuncName, logdata.LineNo, logdata.Message)
}

func (c *ConsoleLogger) Close() {

}
