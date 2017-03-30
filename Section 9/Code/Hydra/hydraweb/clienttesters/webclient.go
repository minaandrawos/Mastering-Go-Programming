package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://requestb.in/1ag9i7m1"
	resp, err := http.Get(url)
	inspectResponse(resp, err)

	data, err := json.Marshal(struct {
		X int
		Y float32
	}{X: 4, Y: 3.8})
	if err != nil {
		log.Fatal("Error occured while marshaling json ", err)
	}

	resp, err = http.Post(url, "application/json", bytes.NewReader(data))
	inspectResponse(resp, err)

	client := http.Client{
		Timeout: 3 * time.Second,
	}
	client.Get(url)

	req, err := http.NewRequest("PUT", "http://requestb.in/rrxka5rr", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("x-testheader", "learning go header")
	req.Header.Set("User-Agent", "Go learning HTTP/1.1")
	resp, err = client.Do(req)
	inspectResponse(resp, err)

	resp, err = http.Get("https://api.ipify.org?format=json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	v := struct {
		IP string `json:"ip"`
	}{}
	err = json.NewDecoder(resp.Body).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(v.IP)
}

func inspectResponse(resp *http.Response, err error) {
	if err != nil {
		log.Fatal("Error occured while marshaling json ", err)
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(b))
}
