package runner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileSrc struct {
	SrcPath     string `json:"srcPath"`     //文件中的原始地址
	OssFullPath string `json:"OssFullPath"` //替换到上传oss后的地址
	OssPath     string `json:"ossPath"`     //上传oss的地址
	LocalPath   string `json:"localPath"`   //本地所处的本地地址
}

func FileGetAll(path string) (files []string, err error) {
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err) // 打印错误信息
			return err       // 可以选择返回错误，但这里我们选择继续遍历
		}
		if !info.IsDir() {
			files = append(files, path) // 将文件路径添加到切片中
		}
		return nil
	})

	fmt.Println(files, err)
	return files, nil
}

func (w *WebSite) ParseFileSrc(unZipPath string) (fileList []FileSrc, err error) {
	unZipPath = strings.ReplaceAll(unZipPath, "\\", "/")
	refPath := unZipPath
	//unZipPath = "./" + unZipPath //./soft/beiluo/json-tool/json_conv/v1.0/dist
	files, err := FileGetAll(unZipPath)
	if err != nil {
		return nil, err
	}
	var filesSrc []FileSrc
	for _, file := range files {
		filePath := strings.ReplaceAll(file, "\\", "/")
		srcPath := strings.ReplaceAll(filePath, refPath, "")
		filesSrc = append(filesSrc, FileSrc{
			SrcPath:     srcPath,
			LocalPath:   "./" + filePath,
			OssFullPath: w.Host + filePath,
			OssPath:     filePath,
		})
	}

	return filesSrc, nil
}

// FileSrcReplaceAndUpload 替换文件中的地址用oss的地址
func (w *WebSite) FileSrcReplaceAndUpload(fileList []FileSrc) error {
	fileReplace := make(map[string]string)
	for _, file := range fileList {
		fileReplace[file.SrcPath] = file.OssFullPath
	}

	for _, src := range fileList {
		fileBytes, err := os.ReadFile(src.LocalPath)
		if err != nil {
			return err
		}
		fileContent := string(fileBytes)
		for k, v := range fileReplace { //把文件中的本地地址替换成oss地址
			fileContent = strings.ReplaceAll(fileContent, k, v)
		}

		err = os.Remove(src.LocalPath)
		if err != nil {
			return err
		}
		create, err := os.Create(src.LocalPath)
		if err != nil {
			return err
		}
		_, err = create.Write([]byte(fileContent))
		if err != nil {
			return err
		}
		create.Close()
	}

	return nil
}

//func (w *WebSite) Deploy() (index string, err error) {
//
//
//}
