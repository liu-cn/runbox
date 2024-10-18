package dto

type GetSendCodeReq struct {
}

type GetSendCodeResp struct {
	ReceiveCode string `json:"receive_code" runner:"desc:取件码;mock_data:2453"`
}
type SenderSendMsgReq struct {
	//ReceiveCode string `json:"receive_code" runner:"desc:取件码;mock_data:2453;required:必填"`
	Content string `json:"content" runner:"desc:内容;mock_data:你好啊;required:必填"`
}

type SenderSendMsgResp struct {
	ReceiveCode string `json:"receive_code" runner:"desc:取件码;mock_data:2453;required:必填"`
	//Content     string `json:"content" runner:"desc:内容;mock_data:你好啊;required:必填"`
}

type SenderReceiveReq struct {
	ReceiveCode string `json:"receive_code" runner:"desc:取件码;mock_data:2453;required:必填"`
}
type SenderReceiveResp struct {
	Content string `json:"content" runner:"desc:内容;mock_data:你好啊;required:必填"`
}
