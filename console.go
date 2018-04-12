package logger

import "os"

type ConsoleLogger struct {
	level int
}

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
	WriteLog(os.Stdout, DebugLevel, format, args...)
}

func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	if c.level > TraceLevel {
		return
	}
	WriteLog(os.Stdout, TraceLevel, format, args...)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	if c.level > InfoLevel {
		return
	}
	WriteLog(os.Stdout, InfoLevel, format, args...)
}

func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	if c.level > WarnLevel {
		return
	}
	WriteLog(os.Stdout, WarnLevel, format, args...)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	if c.level > ErrorLevel {
		return
	}
	WriteLog(os.Stdout, ErrorLevel, format, args...)
}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	if c.level > FatalLevel {
		return
	}
	WriteLog(os.Stdout, FatalLevel, format, args...)
}

func (c *ConsoleLogger) Close() {

}
