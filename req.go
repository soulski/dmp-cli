package dmpc

import (
	"encoding/json"
)

const (
	REQ_RES = "req-res"
	PUB_SUB = "pub-sub"
)

type Service struct {
	Namespace    string `json:"namespace"`
	ContactPoint string `json:"contact-point"`
}

type Message struct {
	Type      string `json:"type"`
	Topic     string `json:"topic"`
	Namespace string `json:"namespace"`
	Body      []byte `json:"body"`
}

func NewReqresMsg(ns string, body interface{}) (*Message, error) {
	bodyBuff, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return &Message{
		Type:      REQ_RES,
		Namespace: ns,
		Body:      bodyBuff,
	}, nil
}

func NewPubsubMsg(topic string, body interface{}) (*Message, error) {
	bodyBuff, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return &Message{
		Type:  PUB_SUB,
		Topic: topic,
		Body:  bodyBuff,
	}, nil
}
