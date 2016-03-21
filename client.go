package dmpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	JSON_CONTENT_TYPE = "application/json"
)

func requestJSON(method string, url string, data interface{}) ([]byte, error) {
	buff, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(buff))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", JSON_CONTENT_TYPE)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func RegisterService(service *Service) (bool, error) {
	_, err := requestJSON("PUT", "http://127.0.0.1:8080/service", service)

	if err != nil {
		return false, err
	}

	return true, err
}

func SendMsg(msg *Message) (*bytes.Reader, error) {
	res, err := requestJSON("PUT", "http://127.0.0.1:8080/message", msg)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(res), err
}

func GetMembers(ns string) (*Members, error) {
	resp, err := http.Get("http://127.0.0.1:8080/member/" + ns)
	if err != nil {
		return nil, err
	}

	var members Members

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&members); err != nil {
		return nil, err
	}

	return &members, nil
}

func GetAllMembers() (*Members, error) {
	resp, err := http.Get("http://127.0.0.1:8080/member")
	if err != nil {
		return nil, err
	}

	var members Members

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&members); err != nil {
		return nil, err
	}

	return &members, nil
}

func SubscribeTopic(topic string) error {
	url := "http://127.0.0.1:8080/topic/%s/subscriber"
	_, err := requestJSON("PUT", fmt.Sprintf(url, topic), "")
	if err != nil {
		return err
	}

	return err
}
