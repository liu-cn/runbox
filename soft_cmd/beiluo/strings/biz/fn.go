package biz

import (
	"github.com/liu-cn/runbox/sdk/runner"
	"strings"
)

func ReplaceAll(ctx *runner.Context) {

	str := ctx.ReqMap()["str"].(string)
	newStr := ctx.ReqMap()["new_str"].(string)
	oldStr := ctx.ReqMap()["old_str"].(string)
	all := strings.ReplaceAll(str, oldStr, newStr)

	ctx.ResponseOkWithJSON(map[string]interface{}{
		"code": 0,
		"msg":  "ok",
		"data": map[string]interface{}{"res": all},
	})
}

func Split(ctx *runner.Context) {
	str := ctx.ReqMap()["str"].(string)
	separator := ctx.ReqMap()["separator"].(string)
	res := strings.Split(str, separator)
	join := strings.Join(res, "\n")
	ctx.ResponseOkWithJSON(map[string]interface{}{
		"code": 0,
		"msg":  "ok",
		"data": map[string]interface{}{
			"splitString": join,
		},
	})
}
