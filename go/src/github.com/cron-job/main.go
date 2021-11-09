package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ThingsPayload struct {
	Data      map[string]interface{} `json:"data"`
	ThingKey  string                 `json:"thing_key"`
	AccessKey string                 `json:"access_key"`
	TimeStamp uint64                 `json:"time_stamp"`
}

func main() {
	for {
		url := "http://65.0.106.100:33001/old_thing_data"
		method := "POST"
		var datafield = make(map[string]interface{})
		datafield["pulsel"] = 1
		datafield["pulseh"] = 2

		now := time.Now()

		count := 1
		tistamp := now.Add(time.Duration(-count) * time.Minute)
		request := ThingsPayload{
			Data:      datafield,
			ThingKey:  "iot_josh_200",
			AccessKey: "78d9fa4teebt5add59ctb86e1a286477cb147392",
			TimeStamp: uint64(tistamp.Unix()),
		}
		reqBody, err := json.Marshal(request)
		if err != nil {
			fmt.Printf("Old Thing : Error while marshaling %s", err.Error())
			return
		}

		requestpayload := bytes.NewReader(reqBody)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, requestpayload)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("accesskey", "78d9fa4teebt5add59ctb86e1a286477cb147391")
		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
		time.Sleep(1 * time.Minute)
	}
}
