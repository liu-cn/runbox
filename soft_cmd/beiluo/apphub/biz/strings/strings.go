package strings

import (
	"fmt"
	"github.com/liu-cn/runbox/sdk/runner"
	"github.com/liu-cn/runbox/soft_cmd/beiluo/apphub/model/dto"
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

func WithStatisticsTextOpt() runner.Option {
	return func(config *runner.Config) {
		config.Request = dto.StatisticsTextReq{}
		config.Response = dto.StatisticsTextResp{}
		config.EnglishName = "strings.keyword.Statistics"
		config.ChineseName = "字符数量分组统计"
		config.Tags = "字符统计"
		config.Classify = "字符串处理"
		config.ApiDesc = "字符数量分组统计"
	}
}
func StatisticsText(ctx *runner.Context) {
	var r dto.StatisticsTextReq
	var resp dto.StatisticsTextResp

	ctx.ShouldBindJSON(&r)
	if r.Stp == "" {
		resp.Data = fmt.Sprintf("%s %v", r.Keywords, strings.Count(r.Content, r.Keywords))
	} else {
		Keywords := strings.Split(r.Keywords, r.Stp)
		for _, keyword := range Keywords {

			resp.Data = resp.Data + fmt.Sprintf("%s %v\n", keyword, strings.Count(r.Content, keyword))
		}
	}
	ctx.OkWithDataJSON(resp)
}

func WithFormatTextToUpperOpt() runner.Option {
	return func(config *runner.Config) {
		config.Request = dto.FormatTextToUpperReq{}
		config.Response = dto.FormatTextToUpperResp{}
		config.ChineseName = "文本转换大写"
		config.Tags = "文本处理"
		config.Classify = "文本处理"
		config.ApiDesc = "文本处理"
	}
}
func FormatTextToUpper(ctx *runner.Context) {
	var req dto.FormatTextToUpperReq
	var resp dto.FormatTextToUpperResp
	ctx.ShouldBindJSON(&req)
	req.Text = strings.ToUpper(req.Text)
	resp.Data = req.Text
	ctx.OkWithDataJSON(resp)
}
