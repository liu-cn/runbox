package table

import (
	"github.com/liu-cn/runbox/pkg/jsonx"
	"github.com/liu-cn/runbox/runner_code/apphub/model/dto"
	"github.com/liu-cn/runbox/sdk/runner"
)

func Demo(ctx *runner.Context) {
	ctx.ResponseOkWithJSON(dto.BaseResponse{
		Code: 0,
		Msg:  "ok",
		Data: map[string]interface{}{
			"user_list": jsonx.Value(`{
	"title":"用户列表",
  "columns": [
    { "prop": "date", "label": "Date", "width": "180" },
    { "prop": "name", "label": "Name", "width": "180" },
    { "prop": "address", "label": "Address" }
  ],
  "data": [
    { "date": "2016-05-02", "name": "John Smith", "address": "北京" },
    { "date": "2016-05-02", "name": "John Smith", "address": "北京" },
    { "date": "2016-05-02", "name": "John Smith", "address": "北京" },
    { "date": "2016-05-02", "name": "John Smith", "address": "北京" },
    { "date": "2016-05-02", "name": "John Smith", "address": "北京" },
    { "date": "2016-05-02", "name": "John Smith", "address": "北京" },
    { "date": "2016-05-02", "name": "John Smith", "address": "北京" }
  ]
}`),
		},
	})
}
