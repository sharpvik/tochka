package tochka

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/sharpvik/tochka/dto"
)

type ErrorResultLog struct {
	Method string          `json:"method"`
	URL    string          `json:"url"`
	Status string          `json:"status"`
	Result dto.ErrorResult `json:"result"`
}

func (rl *ErrorResultLog) From(r *resty.Response) *ErrorResultLog {
	var result dto.ErrorResult

	_ = json.Unmarshal(r.Body(), &result)

	*rl = ErrorResultLog{
		Method: r.Request.Method,
		URL:    r.Request.URL,
		Status: r.Status(),
		Result: result,
	}

	return rl
}

func (rl *ErrorResultLog) String() string {
	data, _ := json.Marshal(rl)
	return string(data)
}

func (rl *ErrorResultLog) Error() string {
	return rl.String()
}
