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

type Msg struct {
	MsgType string          `json:"type"`
	Body    json.RawMessage `json:"body,omitempty"`
}

func NewMsg(mType string, body interface{}) ([]byte, error) {
	bodyBuff, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	msg, err := json.Marshal(&Msg{
		MsgType: mType,
		Body:    bodyBuff,
	})

	return msg, err
}

func NewMsgWithBuff(mType string, buff []byte) ([]byte, error) {
	msg, err := json.Marshal(&Msg{
		MsgType: mType,
		Body:    buff,
	})

	return msg, err
}

func ParseMsg(body []byte) (*Msg, error) {
	var msg *Msg
	err := json.Unmarshal(body, &msg)
	return msg, err
}

func (m *Msg) ParseBody(msgType interface{}) error {
	return json.Unmarshal(m.Body, msgType)
}
