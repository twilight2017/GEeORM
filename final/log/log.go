package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	//log.Lshortfile支持显示文件名和行号
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info ]\033[0m ", log.LstdFlags|log.Lshortfile)
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

//支持设置日志层级： InfoLevel, ErrorLevel, Disabled
const (
	InfoLevel = iota
	Errorlevel
	Disabled
)

// SetLevel controls log level
func SetLevel(level int) {
	//设置层级时要求加上锁支持
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}

	if Errorlevel < level {
		errorLog.SetOutput(ioutil.Discard)
	}

	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
