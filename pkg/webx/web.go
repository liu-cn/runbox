package webx

import (
	"golang.org/x/sync/errgroup"
	"os"
	"path/filepath"
	"strings"
)

type FileWithPath struct {
	AbsolutePath string
	RelativePath string
	ReplacePath  string
	IsIndexFile  bool

	//SavePath     string
}

// ListFiles 返回指定目录下的所有文件及其路径
func ListFiles(startDir string) ([]*FileWithPath, error) {
	var files []*FileWithPath

	err := filepath.Walk(startDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果是文件夹，跳过
		if info.IsDir() {
			return nil
		}

		// 计算相对路径
		relPath, err := filepath.Rel(startDir, path)
		if err != nil {
			return err
		}

		// 获取文件所在的子目录名称
		subDir := filepath.Dir(relPath)
		if subDir == "." {
			subDir = ""
		} else {
			subDir = filepath.Base(subDir) + "/"
		}

		// 拼接子目录名称和文件名
		filePath := filepath.Join(subDir, filepath.Base(relPath))

		// 记录绝对路径和相对路径
		files = append(files, &FileWithPath{
			AbsolutePath: path,
			RelativePath: filePath,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func DistFiles(startDir string) ([]*FileWithPath, error) {
	files, err := ListFiles(startDir)
	if err != nil {
		return nil, err
	}
	startDir = strings.ReplaceAll(startDir, "\\", "/")
	for i, file := range files {
		//files[i].RelativePath = "/" + strings.ReplaceAll(file.RelativePath, "\\", "/")
		files[i].RelativePath = strings.ReplaceAll(file.RelativePath, "\\", "/")
		files[i].AbsolutePath = strings.ReplaceAll(file.AbsolutePath, "\\", "/")
		files[i].RelativePath = strings.ReplaceAll(file.AbsolutePath, startDir, "")

	}
	return files, nil
}

func ReplaceFilePath(webPath string, pathPrefix string) ([]*FileWithPath, error) {
	pathPrefix = strings.ReplaceAll(pathPrefix, "\\", "/")
	pathPrefix = strings.TrimSuffix(pathPrefix, "/")
	fileReplaceMap := make(map[string]string)
	files, err := DistFiles(webPath)
	//return nil, err
	if err != nil {
		return nil, err
	}
	for i := range files {
		if files[i].RelativePath == "/index.html" {
			files[i].IsIndexFile = true
		}
		files[i].ReplacePath = pathPrefix + files[i].RelativePath
		fileReplaceMap[files[i].RelativePath] = files[i].ReplacePath
	}
	var eg errgroup.Group
	for _, file := range files {
		file := file
		eg.Go(func() error {
			fileBytes, err := os.ReadFile(file.AbsolutePath)
			if err != nil {
				return err
			}
			fileContent := string(fileBytes)
			for path, replacePath := range fileReplaceMap {
				fileContent = strings.ReplaceAll(fileContent, "\""+path+"\"", "\""+replacePath+"\"")
			}

			err = os.Remove(file.AbsolutePath)
			if err != nil {
				return err
			}

			create, err := os.Create(file.AbsolutePath)
			if err != nil {
				return err
			}
			defer create.Close()
			_, err = create.Write([]byte(fileContent))
			if err != nil {
				return err
			}
			return nil
		})
	}
	err = eg.Wait()
	if err != nil {
		return nil, err
	}
	return files, nil
}
