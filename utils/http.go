package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpPostRequest(url string, data interface{}, contentType string) (string, error) {
	jsonStr, _ := json.Marshal(data)
	fmt.Println("json:", string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	req.Header.Add("content-type", contentType)
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), nil
}

func HttpGetRequest(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(all), nil
}
