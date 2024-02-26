package user

import (
	"github.com/NeverStopDreamingWang/goi"
	"net/http"
)

func TestView(request *goi.Request) any {

	return goi.Response{
		Status: http.StatusOK,
		Data: goi.Data{
			Status: http.StatusOK,
			Msg:    "ok",
			Data:   nil,
		},
	}
}
