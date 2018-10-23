package utils

import (
	"net/http"

	"github.com/chenjun-git/umbrella-common/render"

	"github.com/chenjun-git/umbrella-sms/common"
)

var (
	// 构造一个通用的JSON Render
	JSON = render.MakeJSON(common.GetMsg)
)

func PlainText(w http.ResponseWriter, r *http.Request, v string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(v))
}
