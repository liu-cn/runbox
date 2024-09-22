package runner

import (
	"fmt"
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/model/request"
	"github.com/liu-cn/runbox/model/response"
	"github.com/liu-cn/runbox/pkg/compress"
	"github.com/liu-cn/runbox/pkg/osx"
	"github.com/liu-cn/runbox/pkg/slicesx"
	"github.com/liu-cn/runbox/pkg/store"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func NewCmd(runner *model.Runner) *Cmd {
	//dir, _ := os.UserHomeDir()
	dir := "./soft_cmd"

	s := model.PcMap[runner.ToolType]
	fullName := runner.AppCode
	if s == "windows" {
		fullName += fullName + ".exe"
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

type Cmd struct {
	InstallInfo
}

// DeCompressPath 解压临时目录
func (c *Cmd) DeCompressPath() string {
	return filepath.Join(c.TempPath, c.User, c.Name)
}

// GetInstallPath  安装目录
func (c *Cmd) GetInstallPath() string {
	return filepath.Join(c.RootPath, c.User, c.Name)
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
	c.InstallInfo.InstallPath = InstallPath

	err = c.Chmod()
	if err != nil {
		return nil, err
	}
	return &c.InstallInfo, nil

}

func (c *Cmd) UnInstall() (*UnInstallInfo, error) {
	return nil, nil
}

func (c *Cmd) Call(r *request.Call) (*response.Call, error) {
	return nil, nil
}
