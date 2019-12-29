package simple

import (
	"io/ioutil"
	"log"
	"os"
	"yichain-go/api"
)

type Logger struct {
	trace   *log.Logger // 记录所有日志
	info    *log.Logger // 重要的信息
	warning *log.Logger // 需要注意的信息
	error   *log.Logger // 非常严重的问题
}

func New(info string, error string) (logger *Logger) {
	logger = &Logger{}

	infoFile, err := os.OpenFile(info, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	errorFile, err := os.OpenFile(error, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logger.trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.info = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.warning = log.New(infoFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.error = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return
}

func (logger *Logger) Trace(content string, replaces ...interface{}) {
	logger.trace.Println(api.Parse(content, replaces))
}

func (logger *Logger) Info(content string, replaces ...interface{}) {
	logger.info.Println(api.Parse(content, replaces))
}

func (logger *Logger) Warning(content string, replaces ...interface{}) {
	logger.warning.Println(api.Parse(content, replaces))
}

func (logger *Logger) Error(content string, replaces ...interface{}) {
	logger.error.Println(api.Parse(content, replaces))
}
