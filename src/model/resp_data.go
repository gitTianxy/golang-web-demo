package model

type RespData struct {
	Status int `json:"status"`
	Msg string `json:"message"`
	Data interface{} `json:"data"`
	Time int64 `json:"time"`
}