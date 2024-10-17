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

type StatisticsTextReq struct {
	Content  string `json:"content" runner:"desc:统计的内容;required:必填;mock_data:zhangsan(张三),lisi(李四),zhangsan,zhangsan"`
	Keywords string `json:"keywords" runner:"desc:统计的所有关键字;required:必填;mock_data:zhangsan,lisi"`
	Stp      string `json:"stp" runner:"desc:关键字分割符;mock_data:,"`
}

type FormatTextToUpperReq struct {
	Text string `json:"text" runner:"desc:要格式化的文本;required:必填;mock_data:ABCd1"`
}

type FormatTextToUpperResp struct {
	Data string `json:"data" runner:"desc:转换大写后的文本;mock_data:ABCD1"`
}

type StatisticsTextResp struct {
	Data string `json:"data" runner:"desc:统计后的数据;mock_data:zhangsan 2 \n lisi 1"`
}
