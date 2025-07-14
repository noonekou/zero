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

func Response(w http.ResponseWriter, resp interface{}, err error, code int) {
	var body Body
	if err != nil {
		body.Code = code
		body.Msg = err.Error()
	} else {
		body.Msg = "success"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}

func Error(w http.ResponseWriter, err error, code int) {
	Response(w, nil, err, code)
}

func Ok(w http.ResponseWriter, data interface{}) {
	Response(w, data, nil, 0)
}
