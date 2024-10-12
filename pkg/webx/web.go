package webx

import (
	"fmt"
	"github.com/liu-cn/runbox/pkg/slicesx"
	"golang.org/x/sync/errgroup"
	"os"
	"path/filepath"
	"strings"
)

type Path struct {
}

type FileWithPath struct {
	AbsolutePath string
	RelativePath string
	ReplacePath  string
	IsIndexFile  bool
	//List[]

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
	files = slicesx.Sort(files, func(a, b *FileWithPath) bool {
		return len(a.RelativePath) < len(b.RelativePath)
	})
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
	for _, file := range files {
		if file.RelativePath == "/index.html" {
			file.IsIndexFile = true
		}
		if strings.Contains(file.RelativePath, "logo.png") {
			fmt.Printf("Found logo.png in %s\n", file.AbsolutePath)
		}
		file.ReplacePath = pathPrefix + file.RelativePath
		fileReplaceMap[file.RelativePath] = file.ReplacePath
	}
	var eg errgroup.Group
	//lk := &sync.Mutex
	for _, file := range files {
		file := file
		eg.Go(func() error {
			fileBytes, err := os.ReadFile(file.AbsolutePath)
			if err != nil {
				return err
			}
			fileContent := string(fileBytes)
			for path, replacePath := range fileReplaceMap {

				old := fmt.Sprintf(`"%s"`, path)
				newPath := fmt.Sprintf(`"%s"`, replacePath)
				fileContent = strings.ReplaceAll(fileContent, old, newPath) //   old:= "/assets/index.html" ->  new:= http://cdn.geeleo.com/+ossPath

				old = fmt.Sprintf(`"%s"`, strings.TrimPrefix(path, "/")) //"/ass"
				newPath = fmt.Sprintf(`"%s"`, strings.ReplaceAll(replacePath, "http://cdn.geeleo.com/", ""))
				fileContent = strings.ReplaceAll(fileContent, old, newPath) //   old:= "assets/index.html" ->  new:= ossPath
				//TODO 这里应该记录替换日志

				//fileContent = strings.ReplaceAll(fileContent, "\""+path+"\"", "\""+replacePath+"\"")
				//fileContent = strings.ReplaceAll(fileContent, "\""+strings.TrimPrefix(path, "/")+"\"", "\""+replacePath+"\"")
				//fileContent = strings.ReplaceAll(fileContent, "\""+strings.TrimPrefix(path, "\\")+"\"", "\""+replacePath+"\"")
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
