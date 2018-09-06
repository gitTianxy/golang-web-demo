package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"model"
	"encoding/json"
)

func deal_err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	url := "http://localhost:8085/items"
	ret, err := http.Get(url)
	deal_err(err)

	defer ret.Body.Close()

	body, err := ioutil.ReadAll(ret.Body)
	deal_err(err)

	var data model.RespData
	err = json.Unmarshal(body, &data)
	deal_err(err)

	fmt.Printf("status:%v, msg:%v, data:%v, time:%v", data.Status, data.Msg, data.Data, data.Time)
}
