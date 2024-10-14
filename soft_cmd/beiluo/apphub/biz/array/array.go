package array

import (
	"github.com/liu-cn/runbox/pkg/jsonx"
	"github.com/liu-cn/runbox/pkg/slicesx"
	"github.com/liu-cn/runbox/sdk/runner"
	"github.com/liu-cn/runbox/soft_cmd/beiluo/apphub/model/dto"
	"strconv"
	"strings"
)

// Diff 计算补集
func Diff(ctx *runner.Context) {

	req := struct {
		Base    string `json:"base"`
		Current string `json:"current"`
		Stp     string `json:"stp"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.ResponseFailParameter()
		return
	}
	bases := strings.Split(req.Base, req.Stp)
	current := strings.Split(req.Current, req.Stp)
	add, remove := slicesx.Diff(bases, current)
	ctx.ResponseOkWithJSON(map[string]interface{}{
		"code": 0,
		"msg":  "ok",
		"data": map[string]interface{}{
			"add":    strings.Join(add, req.Stp),
			"remove": strings.Join(remove, req.Stp),
		},
	})
}

// Split 分割
func Split(ctx *runner.Context) {
	req := struct {
		List string `json:"list"`
		Stp  string `json:"stp"`
		Size string `json:"size"`
	}{}

	ctx.GetLogger().Infof(jsonx.JSONString(req))
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.ResponseFailParameter()
		return
	}
	bases := strings.Split(req.List, req.Stp)
	size, err := strconv.ParseInt(req.Size, 10, 64)
	if err != nil {
		ctx.ResponseFailParameter()
		return
	}
	//current := strings.Split(req.Current, req.Stp)
	slice := slicesx.SplitSlice(bases, int(size))
	s := ""
	for _, v := range slice {
		s += strings.Join(v, req.Stp) + "\n\n"
	}
	ctx.ResponseOkWithJSON(map[string]interface{}{
		"code": 0,
		"msg":  "ok",
		"data": map[string]interface{}{
			"splitList": s,
		},
	})
}

// ComputeIntersection 计算交集
func ComputeIntersection(ctx *runner.Context) {
	req := struct {
		List1 string `json:"list1"`
		List2 string `json:"list2"`
		Stp   string `json:"stp"`
	}{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.ResponseFailParameter()
		return
	}
	list1 := strings.Split(req.List1, req.Stp)
	list2 := strings.Split(req.List2, req.Stp)
	mp1 := make(map[string]struct{})
	for i := 0; i < len(list1); i++ {
		mp1[list1[i]] = struct{}{}
	}

	var sets []string

	for i := 0; i < len(list2); i++ {
		if _, ok := mp1[list2[i]]; ok {
			sets = append(sets, list2[i])
		}
	}
	ctx.ResponseOkWithJSON(map[string]interface{}{
		"code": 0,
		"msg":  "ok",
		"data": map[string]interface{}{
			"intersection": strings.Join(sets, req.Stp),
		},
	})
}

func WithSplitJoinOpt() runner.Option {
	return func(config *runner.Config) {
		config.Request = dto.SplitJoin{}
		config.Response = dto.SplitJoinResp{}
		config.ApiDesc = "字符串二段分割函数"
	}
}

func SplitJoin(ctx *runner.Context) {
	var r dto.SplitJoin
	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		ctx.ResponseFailParameter()
		return
	}
	res := make(map[string]interface{})
	res["res"] = r.Data
	m := dto.BaseResponse{
		Code: 0,
		Msg:  "ok",
		Data: res,
	}
	if r.Data == "" {
		ctx.ResponseOkWithJSON(m)
		return
	}
	i, err := strconv.ParseInt(r.Index, 10, 64)
	if err != nil {
		i = -1
		err = nil
	}

	if r.Stp1 == "" {
		ctx.ResponseOkWithJSON(m)
		return
	}

	split := strings.Split(r.Data, r.Stp1)
	if r.Stp2 == "" {
		res["res"] = strings.Join(split, "\n")
		ctx.ResponseOkWithJSON(m)
		return
	}
	//zhangsan(张三);李四(lisi)
	var vals []string
	for _, v := range split {

		if r.Index == "" {
			stp2 := strings.Split(v, r.Stp2)
			vals = append(vals, stp2...)

		} else {
			if i != -1 {
				if i == 0 {
					i = 1
				}
				vals = append(vals, strings.Split(v, r.Stp2)[i-1])
			}

		}
	}
	data := strings.Join(vals, r.Stp1)
	res["res"] = data
	ctx.ResponseOkWithJSON(m)
}
