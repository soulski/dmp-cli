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

func request(method string, url string, data []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
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

func RegisterService(service *Service) (*Member, error) {
	rBody, err := requestJSON("PUT", "http://127.0.0.1:8080/namespace", service)

	if err != nil {
		return nil, err
	}

	var member *Member
	err = json.Unmarshal(rBody, &member)
	return member, err
}

func Request(ns string, msg []byte) ([]byte, error) {
	res, err := request("PUT", fmt.Sprintf("http://127.0.0.1:8080/message/reqRes/%s", ns), msg)
	if err != nil {
		return nil, err
	}

	return res, err
}

func Publish(topic string, msg []byte) ([]byte, error) {
	res, err := request("PUT", fmt.Sprintf("http://127.0.0.1:8080/message/pubSub/%s", topic), msg)
	if err != nil {
		return nil, err
	}

	return res, err
}

func Notificate(ns string, msg []byte) ([]byte, error) {
	res, err := request("PUT", fmt.Sprintf("http://127.0.0.1:8080/message/noti/%s", ns), msg)
	if err != nil {
		return nil, err
	}

	return res, err
}

func GetMembers(ns string) (*Members, error) {
	resp, err := http.Get("http://127.0.0.1:8080/namespace/" + ns)
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
	resp, err := http.Get("http://127.0.0.1:8080/namespace")
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
