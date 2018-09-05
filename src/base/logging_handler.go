package base

import (
	"net/http"
	"log"
)

func LogRequest(r *http.Request, msg string)  {
	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		msg,
	)
}

func LogRequestFinish(r *http.Request, msg string, costSecs float64)  {
	log.Printf(
		"%s\t%s\t%s\tcosts:%vsecs",
		r.Method,
		r.RequestURI,
		msg,
		costSecs,
	)
}

func LogRequestErr(r *http.Request, err error)  {
	log.Printf(
		"%s\t%s\t%v",
		r.Method,
		r.RequestURI,
		err,
	)
}