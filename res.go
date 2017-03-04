package dmpc

import (
	"encoding/json"
)

type Members struct {
	Members []*Member `json:"members"`
}

type Member struct {
	IP        string `json:"ip"`
	Status    string `json:"status"`
	Namespace string `json:"namespace"`
}

type Error struct {
	Message string
	Cause   string
}

type Result struct {
	Action  bool            `json:"action"`
	Error   *Error          `json:"error"`
	Message json.RawMessage `json:"message,omitempty"`
}

func CreateMsgResult(msg interface{}) ([]byte, error) {
	buff, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	result := &Result{
		Action:  true,
		Message: buff,
	}

	return json.Marshal(result)
}

func CreateErrorResult(err *Error) ([]byte, error) {
	result := &Result{
		Action: false,
		Error:  err,
	}

	return json.Marshal(result)
}

func ParseResult(buff []byte) (*Result, error) {
	var result *Result
	err := json.Unmarshal(buff, &result)
	return result, err
}

func (r *Result) ParseMsg(msgType interface{}) error {
	return json.Unmarshal(r.Message, msgType)
}
