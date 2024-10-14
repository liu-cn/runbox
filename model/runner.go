package model

import (
	"fmt"
)

type Runner struct {
	StoreRoot  string `json:"store_root"`                                                                          //oss 存储的跟路径
	AppCode    string `json:"app_code" form:"appCode" gorm:"column:app_code;comment:;" binding:"required"`         //应用名称（英文标识）
	ToolType   string `json:"tool_type" form:"tool_type" gorm:"column:tool_type;comment:工具类型;" binding:"required"` //工具类型
	Version    string `json:"version" form:"version" gorm:"column:version;comment:;" binding:"required"`           //应用版本
	OssPath    string `json:"oss_path" form:"ossPath" gorm:"column:oss_path;comment:;"`                            //文件地址
	TenantUser string `json:"tenant_user" form:"tenant_user" gorm:"column:tenant_user;comment:所属租户;"`              //所属租户
}

func (r *Runner) Check() error {
	if r.AppCode == "" {
		return fmt.Errorf("AppCode 不能为空")
	}
	if r.ToolType == "" {
		return fmt.Errorf("ToolType 不能为空")
	}
	if r.Version == "" {
		return fmt.Errorf("version 不能为空")
	}

	if r.TenantUser == "" {
		return fmt.Errorf("TenantUser 不能为空")
	}
	return nil
}

type UpdateVersion struct {
	RunnerConf *Runner `json:"runner_conf"`
	OldVersion string  `json:"old_version"`
	//NewVersion        string  `json:"new_version"`
	//NewVersionOssPath string  `json:"new_version_oss_path"`
}
