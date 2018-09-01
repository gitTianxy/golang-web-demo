package base

import (
	"net/http"
	"encoding/json"
)

type HttpRequestHandler struct {
	Req *http.Request
}

func (handler HttpRequestHandler) GetReqParam(key string) (string, bool) {
	params := handler.Req.URL.Query()
	vals, ok := params[key]
	if !ok {
		return "null", ok
	}
	return vals[0], ok
}

func (handler HttpRequestHandler) GetReqParams(key string) ([]string, bool) {
	params := handler.Req.URL.Query()
	vals, ok := params[key]
	return vals, ok
}

type HttpResponseHandler struct {
	Writer http.ResponseWriter
}

func (h HttpResponseHandler) Succ(msg string) {
	h.Writer.WriteHeader(http.StatusOK)
	h.Writer.Write([]byte(msg))
}

func (h HttpResponseHandler) Fail(msg string)  {
	h.Writer.WriteHeader(http.StatusBadRequest)
	h.Writer.Write([]byte(msg))
}

func (h HttpResponseHandler) NotFound(msg string) {
	h.Writer.WriteHeader(http.StatusNotFound)
	h.Writer.Write([]byte(msg))
}

func (h HttpResponseHandler) HandleResult(ret interface{}) {
	h.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	h.Writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(h.Writer).Encode(ret); err != nil {
		panic(err)
	}
}
