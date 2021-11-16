package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type ThingsPayload struct {
	Data      map[string]interface{} `json:"data"`
	ThingKey  string                 `json:"thing_key"`
	AccessKey string                 `json:"access_key"`
	TimeStamp uint64                 `json:"time_stamp"`
}

var valpul = 1
var valpolh = 2
var valwat = 20
var valval = 30

func dataSeed(key string) {
	url := "http://65.0.106.100:33001/old_thing_data"
	method := "POST"
	var datafield = make(map[string]interface{})
	datafield["pulsel"] = valpul
	datafield["pulseh"] = valpolh
	datafield["meter_status"] = valwat
	datafield["meter_status"] = valval
	now := time.Now()

	count := 1
	tistamp := now.Add(time.Duration(-count) * time.Minute)
	request := ThingsPayload{
		Data:      datafield,
		ThingKey:  key,
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
	valpul += 1
	valpolh += 1
	valwat += 1
	valval += 1
}

func main() {
	file, err := os.Open("thing_key.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, val := range records[0] {
		dataSeed(val)
	}
}
