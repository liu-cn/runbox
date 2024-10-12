package runner

import (
	"fmt"
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/model/request"
	"github.com/liu-cn/runbox/model/response"
	"github.com/liu-cn/runbox/pkg/compress"
	"github.com/liu-cn/runbox/pkg/osx"
	"github.com/liu-cn/runbox/pkg/store"
	"github.com/liu-cn/runbox/pkg/webx"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func NewWebSite(runner *model.Runner) *WebSite {
	dir := "./web"
	fullName := runner.AppCode
	return &WebSite{
		Host: "http://cdn.geeleo.com",
		InstallInfo: response.InstallInfo{
			TempPath:     filepath.Join(os.TempDir(), runner.ToolType),
			RootPath:     dir,
			StoreRoot:    runner.StoreRoot,
			Name:         runner.AppCode,
			FullName:     fullName,
			User:         runner.TenantUser,
			Version:      runner.Version,
			DownloadPath: runner.OssPath,
		},
	}
}

type WebSite struct {
	Host string
	response.InstallInfo
}

// DeCompressPath 解压临时目录
func (w *WebSite) DeCompressPath() string {
	return filepath.Join(w.TempPath, w.User, w.Name)
}

func findIndexFile(files []*webx.FileWithPath) *webx.FileWithPath {
	for i, file := range files {
		if file.IsIndexFile {
			return files[i]
		}
	}
	return nil
}

func backIndexFile(files []*webx.FileWithPath) *webx.FileWithPath {
	for _, file := range files {
		if file.IsIndexFile {
			b := *file
			b.IsIndexFile = false
			return &b
		}
	}
	return nil
}

func (w *WebSite) GetSavePath(path *webx.FileWithPath) (savePath string) {
	path.RelativePath = strings.TrimPrefix(path.RelativePath, "/")
	path.RelativePath = strings.TrimPrefix(path.RelativePath, "\\")
	if path.IsIndexFile {
		savePath = fmt.Sprintf("%s/%s/%s/%s",
			w.StoreRoot, w.User, w.Name,
			strings.ReplaceAll(path.RelativePath, ".html", "")) //index.html
	} else {
		savePath = fmt.Sprintf("%s/%s/%s/%s/%s",
			w.StoreRoot, w.User, w.Name, w.Version,
			path.RelativePath)
	}

	return savePath
}

func (w *WebSite) Install(store store.FileStore) (installInfo *response.InstallInfo, err error) {

	file, err := store.GetFile(w.DownloadPath)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(file.FileLocalPath)
	DeCompressPath := w.DeCompressPath()

	err = compress.DeCompress(file.FileLocalPath, DeCompressPath)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(DeCompressPath)
	fileList, dirs, err := osx.CheckDirectChildren(DeCompressPath) //检查解压后的文件目录，如果只有一个文件夹，说明文件在下一级目录
	if err != nil {
		return nil, err
	}
	rootPath := DeCompressPath
	if len(fileList) == 0 && len(dirs) == 1 {
		rootPath = filepath.Join(rootPath, dirs[0]) //取下级目录
	}
	now := time.Now()
	files, err := webx.ReplaceFilePath(rootPath, fmt.Sprintf("%s/%s/%s/%s/%s",
		w.Host, w.StoreRoot, w.User, w.Name, w.Version))

	if err != nil {
		return nil, err
	}
	cost1 := time.Since(now)

	indexFile := findIndexFile(files)
	if indexFile == nil {
		return nil, fmt.Errorf("index.html file not found")
	}
	backIndex := backIndexFile(files)
	files = append(files, backIndex)
	wg := &sync.WaitGroup{}
	wg.Add(len(files))
	t := time.Now()
	for _, fileInfo := range files {
		fileInfo := fileInfo
		go func() {
			defer wg.Done()
			savePath := w.GetSavePath(fileInfo)
			fmt.Println(savePath)
			if fileInfo.IsIndexFile {
				del := savePath
				err := store.DeleteFile(del)
				if err != nil {
					fmt.Println(err.Error())
					//return nil, err
				}
			}
			if strings.Contains(fileInfo.AbsolutePath, "logo.png") {
				fmt.Println(fileInfo)
			}
			_, err := store.FileSave(fileInfo.AbsolutePath, w.GetSavePath(fileInfo))
			if err != nil {
				fmt.Println(fmt.Sprintf("%+v err:%s", fileInfo, err.Error()))
				//return nil, err
			}

		}()
	}
	wg.Wait()
	fmt.Println("上传文件耗时：", time.Since(t))
	fmt.Println("解析文件耗时：", cost1)
	return &response.InstallInfo{}, nil
}
func (w *WebSite) RollbackVersion(r *request.RollbackVersion, fileStore store.FileStore) (*response.RollbackVersion, error) {
	return nil, nil
}

func (w *WebSite) GetInstallPath() (path string) {
	//TODO implement me
	panic("implement me")
}

func (w *WebSite) UnInstall() (unInstallInfo *response.UnInstallInfo, err error) {
	//TODO implement me
	panic("implement me")
}

func (w *WebSite) UpdateVersion(up *model.UpdateVersion, fileStore store.FileStore) (*response.UpdateVersion, error) {
	_, err := w.Install(fileStore)
	if err != nil {
		return nil, err
	}
	return &response.UpdateVersion{}, nil
}

func (w *WebSite) Run(req *request.Run) (*response.Run, error) {
	//TODO implement me
	panic("implement me")
}
