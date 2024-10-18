package sender

import (
	"errors"
	"fmt"
	"github.com/liu-cn/runbox/sdk/runner"
	"github.com/liu-cn/runbox/soft_cmd/beiluo/apphub/model/dto"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Base struct {
	Id        int64          `gorm:"primaryKey;autoIncrement"`
	User      string         `json:"user"`
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;autoUpdatedTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
type SendRecord struct {
	Base
	Received    int64  `json:"received" gorm:"column:received;comment:是否已经被接收"`
	ReceiveCode string `json:"receive_code" gorm:"column:receive_code;comment:取件码"`
	Content     string `json:"content" gorm:"column:content;type:longtext;comment:消息内容"`
	MsgType     string `json:"msg_type;comment:消息类型"`
}

func (s *SendRecord) TableName() string {
	return "send_record"
}

var senderDb *gorm.DB

func generateRandomNumber() string {
	rand.Seed(time.Now().UnixNano()) // 用当前时间作为随机数生成的种子
	num := rand.Intn(9000) + 1000    // 生成一个1000到9999之间的随机数
	return fmt.Sprintf("%04d", num)  // 格式化为4位数
}
func getSenderDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sender.db"))
	if err != nil {
		return nil
	}
	return db
}

func genCode() string {
	var codes []string
	mp := make(map[string]bool)
	for range 10 {
		code := generateRandomNumber()
		mp[code] = true
		codes = append(codes, code)
	}
	var records []SendRecord
	err := getSenderDB().Model(&SendRecord{}).Where("receive_code in ? AND received=0", codes).Find(&records).Error
	if err != nil {
		return generateRandomNumber() + "2"
	}
	if len(records) == 0 {
		return codes[0]
	}
	mps := make(map[string]bool)
	for _, record := range records {
		mps[record.ReceiveCode] = true
	}
	for _, code := range codes {
		if !mps[code] {
			return code
		}
	}
	return generateRandomNumber() + "1"

}

func WithGetSendCodeOpt() runner.Option {
	return func(config *runner.Config) {
		//config.Request = dto.StatisticsTextReq{}
		config.Response = dto.GetSendCodeReq{}
		config.EnglishName = "sender.genReceiveCode"
		config.ChineseName = "文件传输-生成取件码"
		config.Tags = "文件传输"
		config.Classify = "文件传输"
		config.ApiDesc = "文件传输-生成取件码，生成的取件码可以取出发送的文件或者文本"
	}
}

//	func GetSendCode(ctx *runner.Context) {
//		//getSenderDB().Create(&model.Sender{})
//		ctx.ResponseOkWithJSON(dto.GetSendCodeResp{
//			Code: genCode(),
//		})
//	}
func WithSendMsgOpt() runner.Option {
	return func(config *runner.Config) {
		config.Request = dto.SenderSendMsgReq{}
		config.Response = dto.SenderSendMsgResp{}
		config.EnglishName = "message.Send"
		config.ChineseName = "文件传输-发送文本消息"
		config.Tags = "文件传输"
		config.Classify = "文件传输"
		config.ApiDesc = "文件传输-发送文本消息，返回取件码，根据取件码可以取出消息"
	}
}
func SendMsg(ctx *runner.Context) {
	getSenderDB().AutoMigrate(&SendRecord{})
	//getSenderDB().Create(&model.Sender{})
	var req dto.SenderSendMsgReq
	var res dto.SenderSendMsgResp
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.ResponseFailParameter()
		return
	}
	code := genCode()
	err = getSenderDB().Model(&SendRecord{}).Create(&SendRecord{
		ReceiveCode: code,
		Content:     req.Content,
		Received:    0,
		MsgType:     "text",
	}).Error
	if err != nil {
		ctx.ResponseFailDefaultJSONWithMsg("内部错误！")
		return
	}
	res.ReceiveCode = code
	ctx.OkWithDataJSON(res)
}

func WithReceiveMsgOpt() runner.Option {
	return func(config *runner.Config) {
		config.Request = dto.SenderReceiveReq{}
		config.Response = dto.SenderReceiveResp{}
		config.EnglishName = "message.Receive"
		config.ChineseName = "文件传输-接收消息"
		config.Tags = "文件传输"
		config.Classify = "文件传输"
		config.ApiDesc = "文件传输-接收消息，输入取件码取出文本消息"
	}
}
func ReceiveMsg(ctx *runner.Context) {
	//getSenderDB().Create(&model.Sender{})
	var req dto.SenderReceiveReq
	var res dto.SenderReceiveResp
	ctx.ShouldBindJSON(&req)
	r := SendRecord{}

	err := getSenderDB().Model(&SendRecord{}).Where("receive_code = ? AND received = ?", req.ReceiveCode, 0).First(&r).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.ResponseFailDefaultJSONWithMsg("已经取件或者取件码错误！")
			return
		}
		ctx.ResponseFailDefaultJSONWithMsg("内部错误！" + err.Error())
		return
	}
	if r.Content == "" {
		return
	}
	if r.Content != "" {
		err = getSenderDB().Model(&SendRecord{}).Where("receive_code = ?", req.ReceiveCode).
			UpdateColumn("received", 1).Error

	}
	res.Content = r.Content
	ctx.OkWithDataJSON(res)
}
