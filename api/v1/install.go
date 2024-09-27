package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/model/response"
	"github.com/sirupsen/logrus"
)

// Install 安装软件
func (r *Api) Install(c *gin.Context) {
	var (
		runnerModel model.Runner
		err         error
	)
	defer func() {
		logrus.Infof("Install req:%+v err:%v", runnerModel, err)
	}()
	err = c.ShouldBindJSON(&runnerModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	install, err := r.RunBox.Install(&runnerModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(install, c)
}

// UpdateVersion 更新版本
func (r *Api) UpdateVersion(c *gin.Context) {
	var (
		runnerModel model.UpdateVersion
		err         error
	)
	defer func() {
		logrus.Infof("UpdateVersion req:%+v err:%v", runnerModel, err)
	}()
	err = c.ShouldBindJSON(&runnerModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	install, err := r.RunBox.UpdateVersion(&runnerModel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(install, c)
}
