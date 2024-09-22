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

// InstallPath 解压临时目录
func (c *Cmd) InstallPath() string {
	return filepath.Join(c.RootPath, c.User, c.Name)
}

func (c *Cmd) Chmod() error {
	if runtime.GOOS != "windows" {
		cmdPath := filepath.Join(c.InstallPath(), c.Name)
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
	InstallPath := c.InstallPath()
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
	////absPath = strings.Join([]string{absPath, c.User, c.Name}, "/")
	//absPath = strings.Join([]string{absPath, c.User}, "/")
	//unZipOut := absPath + "/" + c.Name
	//path := strings.Split(c.DownloadPath, "/")
	//fileName := path[len(path)-1]
	//
	//appName := ""
	//if runtime.GOOS == "windows" {
	//	//fileName =soft.zip
	//	appName = strings.Split(fileName, ".")[0] + ".exe"
	//} else {
	//	appName = strings.Split(fileName, ".")[0]
	//}
	//c.FullName = appName
	//out := absPath + "/" + fileName
	//defer os.Remove(out)
	//os.MkdirAll(absPath, os.ModePerm)
	//url := "http://cdn.geeleo.com/" + c.DownloadPath
	//err := httpx.DownloadFile(url, out)
	//if err != nil {
	//	return nil, err
	//}
	//
	//unZipPath, err := compress.UnZip(filepath.Join(absPath, fileName), unZipOut)
	//if err != nil {
	//	return nil, err
	//}
	////todo 设置权限
	////exec.Command("chmod")
	//// 创建一个命令来添加执行权限
	//if runtime.GOOS != "windows" {
	//	p := unZipOut + "/" + c.Name
	//	cmd := exec.Command("chmod", "+x", unZipOut+"/"+c.Name)
	//	// 执行命令
	//	err = cmd.Run()
	//	if err != nil {
	//		fmt.Printf("cmd.Run() failed with p:%s err:%s\n", p, err)
	//		return nil, err
	//	}
	//}
	//c.InstallPath = unZipPath
	////判断是否存在该软件
	//return &c.InstallInfo, nil
}

//func (c *Cmd) Install() (*InstallInfo, error) {
//	absPath := "./soft_cmd"
//	//absPath = strings.Join([]string{absPath, s.User, s.Name}, "/")
//	absPath = strings.Join([]string{absPath, c.User}, "/")
//	unZipOut := absPath + "/" + c.Name
//	path := strings.Split(c.DownloadPath, "/")
//	fileName := path[len(path)-1]
//
//	appName := ""
//	if runtime.GOOS == "windows" {
//		//fileName  =soft.zip
//		appName = strings.Split(fileName, ".")[0] + ".exe"
//	} else {
//		appName = strings.Split(fileName, ".")[0]
//	}
//	c.FullName = appName
//	out := absPath + "/" + fileName
//	defer os.Remove(out)
//	os.MkdirAll(absPath, os.ModePerm)
//	url := "http://cdn.geeleo.com/" + c.DownloadPath
//	err := httpx.DownloadFile(url, out)
//	if err != nil {
//		return nil, err
//	}
//
//	unZipPath, err := compress.UnZip(filepath.Join(absPath, fileName), unZipOut)
//	if err != nil {
//		return nil, err
//	}
//	//todo 设置权限
//	//exec.Command("chmod")
//	// 创建一个命令来添加执行权限
//	if runtime.GOOS != "windows" {
//		p := unZipOut + "/" + c.Name
//		cmd := exec.Command("chmod", "+x", unZipOut+"/"+c.Name)
//		// 执行命令
//		err = cmd.Run()
//		if err != nil {
//			fmt.Printf("cmd.Run() failed with p:%s err:%s\n", p, err)
//			return nil, err
//		}
//	}
//	c.InstallPath = unZipPath
//	//判断是否存在该软件
//	return &c.InstallInfo, nil
//}

func (c *Cmd) UnInstall() (*UnInstallInfo, error) {
	return nil, nil
}

func (c *Cmd) Call(r *request.Call) (*response.Call, error) {
	return nil, nil
}
