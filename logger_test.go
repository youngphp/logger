package logger

import "testing"

func TestFileLogger(t *testing.T) {
	log := NewFileLogger(DebugLevel, "f:/testlog")
	log.Debug("the use is %s age is %d", "fei", 20)
	log.Warn("the warn is %s", "open file is error")
	log.Fatal("program is die %s", "程序崩溃了")

	log = NewConsoleLogger(DebugLevel)
	log.Debug("the use is %s age is %d", "fei", 20)
	log.Warn("the warn is %s", "open file is error")
	log.Fatal("program is die %s", "程序崩溃了")
}
