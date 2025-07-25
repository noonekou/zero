package response

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = -1
		body.Msg = err.Error()
	} else {
		body.Msg = "success"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}

func ResponseError(w http.ResponseWriter, err error) {
	Response(w, nil, err)
}

func ResponseOk(w http.ResponseWriter, data interface{}) {
	Response(w, data, nil)
}
