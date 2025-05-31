package dto

import (
	"encoding/json"
	"errors"
)

type ErrorResult struct {
	Code    string  `json:"code"`
	ID      string  `json:"id"`
	Message string  `json:"message"`
	Errors  []Error `json:"Errors"`
}

type Error struct {
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	URL       string `json:"url"`
}

func (err *ErrorResult) From(body []byte) error {
	var result ErrorResult
	if err := json.Unmarshal(body, &result); err != nil {
		return errors.New(string(body))
	}

	*err = result

	return err
}

func (err *ErrorResult) Error() string {
	data, _ := json.Marshal(err)
	return string(data)
}
