package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	//使用log.Lshortfile支持显示文件名和代码行号
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os, Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog}
	mu       sync.Mutex
)

//log methods
var (
	Error  = errorLog.Println
	Errorf = errorLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

// log levels
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

//Setlevel controls log level
func Setlevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutPut(os.Stdout)
	}
	if ErrorLevel < level {
		errorLog.SetOutPut(ioutil.Discard)
	}

	if InfoLevel < level {
		infolog.SetOutPut(ioutil.Discard)
	}
}
