package base

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type HttpRequestHandler struct {
	Req *http.Request
}

func (handler HttpRequestHandler) GetParamVal(key string) string {
	params := handler.Req.URL.Query()
	vals, ok := params[key]
	if !ok {
		panic(fmt.Sprintf("invalid param: %v", key))
	}
	return vals[0]
}

type HttpResponseHandler struct {
	Writer http.ResponseWriter
}

func (handler HttpResponseHandler) Succ() {
	handler.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	handler.Writer.WriteHeader(http.StatusOK)
}

func (handler HttpResponseHandler) NotFound() {
	handler.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	handler.Writer.WriteHeader(http.StatusNotFound)
}

func (handler HttpResponseHandler) HandleResult(ret interface{}) {
	handler.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	handler.Writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(handler.Writer).Encode(ret); err != nil {
		panic(err)
	}
}
