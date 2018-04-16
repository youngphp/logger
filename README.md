## :sweat_drops:golang版的日志调试工具
***
## how to use 
    go get -u github.com/youngphp/logger  
***
### 详情可查看logger_test.go
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