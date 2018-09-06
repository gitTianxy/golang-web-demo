package base

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const SERVER_LOG_PATH = "log/server.log"

func getLogFile(fpath string) (f *os.File, err error)  {
	f, err = os.OpenFile(fpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	return
}

func GetDebugLogger() *log.Logger  {
	return log.New(os.Stdout, "[DEBUG]", log.Ldate|log.Ltime)
}

func GetInfoLogger() *log.Logger {
	f, _ := getLogFile(SERVER_LOG_PATH)
	return log.New(io.MultiWriter(os.Stdout, f), "[INFO]", log.Ldate|log.Ltime)
}

func GetWarnLogger() *log.Logger {
	f, _ := getLogFile(SERVER_LOG_PATH)
	return log.New(io.MultiWriter(os.Stdout, f), "[WARN]", log.Ldate|log.Ltime)
}

func GetErrorLogger() *log.Logger {
	f, _ := getLogFile(SERVER_LOG_PATH)
	return log.New(io.MultiWriter(os.Stdout, f), "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
}

type RestfulLogHandler struct {
	infoLogger *log.Logger
	errorLogger *log.Logger
}

func GetRestfulLogger() *RestfulLogHandler  {
	h := RestfulLogHandler{}
	h.infoLogger = GetInfoLogger()
	h.errorLogger = GetErrorLogger()
	return &h
}

func (h RestfulLogHandler) LogRequest(r *http.Request, msg string) {
	h.infoLogger.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		msg,
	)
}

func (h RestfulLogHandler) LogRequestFinish(r *http.Request, msg string, duration time.Duration) {
	h.infoLogger.Printf(
		"%s\t%s\t%s\tcosts:%vsecs",
		r.Method,
		r.RequestURI,
		msg,
		duration.Seconds(),
	)
}

func (h RestfulLogHandler) LogRequestErr(r *http.Request, err error) {
	h.errorLogger.Printf(
		"%s\t%s\t%v",
		r.Method,
		r.RequestURI,
		err,
	)
}
