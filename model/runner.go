package model

type Runner struct {
	AppCode    string `json:"appCode" form:"appCode" gorm:"column:app_code;comment:;" binding:"required"`          //应用名称（英文标识）
	ToolType   string `json:"tool_type" form:"tool_type" gorm:"column:tool_type;comment:工具类型;" binding:"required"` //工具类型
	Version    string `json:"version" form:"version" gorm:"column:version;comment:;" binding:"required"`           //应用版本
	OssPath    string `json:"ossPath" form:"ossPath" gorm:"column:oss_path;comment:;"`                             //文件地址
	TenantUser string `json:"tenant_user" form:"tenant_user" gorm:"column:tenant_user;comment:所属租户;"`              //所属租户
}
