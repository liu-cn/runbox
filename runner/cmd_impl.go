package runner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/model/request"
	"github.com/liu-cn/runbox/model/response"
	"github.com/liu-cn/runbox/pkg/compress"
	"github.com/liu-cn/runbox/pkg/osx"
	"github.com/liu-cn/runbox/pkg/slicesx"
	"github.com/liu-cn/runbox/pkg/store"
	"github.com/liu-cn/runbox/pkg/stringsx"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func NewCmd(runner *model.Runner) *Cmd {
	//dir, _ := os.UserHomeDir()
	dir := "./soft_cmd"

	//s := model.PcMap[runner.ToolType]
	fullName := runner.AppCode
	if runner.ToolType == "windows" {
		fullName += ".exe"
	}
	return &Cmd{
		InstallInfo: InstallInfo{
			TempPath:     os.TempDir(),
			RootPath:     dir,
			Name:         runner.AppCode,
			FullName:     fullName,
			User:         runner.TenantUser,
			Version:      runner.Version,
			DownloadPath: runner.OssPath,
		},
	}
}

func (c *Cmd) GetAppName() string {
	return c.FullName
}

type Cmd struct {
	InstallInfo
}

// DeCompressPath 解压临时目录
func (c *Cmd) DeCompressPath() string {
	return filepath.Join(c.TempPath, c.User, c.Name)
}

// GetInstallPath  安装目录
func (c *Cmd) GetInstallPath() string {
	abs, err := filepath.Abs(fmt.Sprintf("%s/%s/%s", c.RootPath, c.User, c.Name))
	if err != nil {
		panic(err)
		return abs
	}
	return abs
}

func (c *Cmd) Chmod() error {
	if runtime.GOOS != "windows" {
		cmdPath := filepath.Join(c.GetInstallPath(), c.Name)
		cmd := exec.Command("chmod", "+x", cmdPath)
		// 执行命令
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("Chmod cmd.Run() failed with path:%s err:%s\n", cmdPath, err)
		}
	}
	return nil
}

func (c *Cmd) Install(fileStore store.FileStore) (*InstallInfo, error) {
	//absPath := c.RootPath

	file, err := fileStore.GetFile(c.DownloadPath)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(file.FileLocalPath)

	DeCompressPath := c.DeCompressPath()
	defer os.RemoveAll(DeCompressPath)
	err = compress.DeCompressx(file.FileLocalPath, DeCompressPath)
	if err != nil {
		return nil, err
	}
	files, dirs, err := osx.CheckDirectChildren(DeCompressPath)
	if err != nil {
		return nil, err
	}
	if len(files) == 0 && len(dirs) == 0 {
		return nil, fmt.Errorf("cmd Install no files found in %s", DeCompressPath)
	}
	dirs = slicesx.Select(dirs, func(t string) bool {
		return !strings.HasPrefix(t, "_")
	})
	compressHome := ""
	if len(files) == 0 && len(dirs) == 1 {
		if dirs[0] == c.Name { //说明解压时候解压路径多出一级和应用名称相同的目录
			compressHome = filepath.Join(DeCompressPath, c.Name)
		}
	} else {
		compressHome = DeCompressPath
	}
	InstallPath := c.GetInstallPath()
	err = osx.CopyDirectory(compressHome, InstallPath)
	if err != nil {
		return nil, err
	}
	//c.InstallInfo.InstallPath = InstallPath

	err = c.Chmod()
	if err != nil {
		return nil, err
	}
	return &c.InstallInfo, nil

}

func (c *Cmd) UnInstall() (*UnInstallInfo, error) {
	return nil, nil
}

func (c *Cmd) Call(req *request.Run) (*response.Run, error) {
	now := time.Now()
	installPath := c.GetInstallPath()
	appName := c.GetAppName()
	softPath := fmt.Sprintf("%s/%s", installPath, appName)
	softPath = strings.ReplaceAll(softPath, "\\", "/")
	req.RequestJsonPath = strings.ReplaceAll(req.RequestJsonPath, "\\", "/")
	cmd := exec.Command(softPath, req.Command, req.RequestJsonPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	cmdStr := fmt.Sprintf("%s %s %s", softPath, req.Command, req.RequestJsonPath)
	fmt.Println(cmdStr)
	if err != nil {
		return nil, err
	}
	s := out.String()
	if s == "" {
		//todo
		return nil, fmt.Errorf("out.String() ==== nil cmd程序输出的结果为空，请检测程序是否正确")
	}
	resList := stringsx.ParserHtmlTagContent(s, "Response")
	if len(resList) == 0 {
		//todo 请使用sdk开发软件
		return nil, fmt.Errorf("soft call err 请使用sdk开发软件")
	}
	var res response.Run
	err = json.Unmarshal([]byte(resList[0]), &res)
	if err != nil {
		return nil, err
	}
	since := time.Since(now)
	res.CallCostTime = since
	res.ResponseMetaData = s
	//p.printSoftLogs(s, since)
	return &res, nil
}
