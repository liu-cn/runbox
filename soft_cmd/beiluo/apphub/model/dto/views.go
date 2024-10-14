package dto

type ViewReq struct {
	FilePath string `json:"file_path"`
}

type SplitJoin struct {
	Data  string `json:"data" runner:"desc:原始数据;required:必填;mock_data:zhangsan(张三),lisi(李四)"`
	Stp1  string `json:"stp1" runner:"desc:一段分隔符;mock_data:,"`
	Stp2  string `json:"stp2" runner:"desc:二段分隔符;mock_data:("`
	Index string `json:"index" runner:"desc:二段分后取索引位置;mock_data:1"`
}

type SplitJoinResp struct {
	Res string `json:"res" runner:"desc:处理后数据;mock_data:zhangsan,lisi"`
}

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
