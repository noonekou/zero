package response

import (
	errs "bookstore/common/error"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		if myErr := errs.GetMyError(err); myErr != nil {
			body.Code = int(myErr.Code)
			body.Msg = myErr.Message
		} else {
			if st := status.Convert(err); st != nil {
				body.Code = int(st.Code())
				body.Msg = st.Message()
			} else {
				body.Code = int(errs.ErrCodeFail)
				body.Msg = err.Error()
			}
		}
	} else {
		body.Msg = "success"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}

func Error(w http.ResponseWriter, err error) {
	Response(w, nil, err)
}

func Ok(w http.ResponseWriter, data interface{}) {
	Response(w, data, nil)
}
