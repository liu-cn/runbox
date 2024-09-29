package webx

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestReplaceFilePath(t *testing.T) {
	files, err := ListFiles("D:\\code\\github.com\\apphub\\web\\dist")
	if err != nil {
		panic(err)
	}
	fileList := []string{}
	for _, file := range files {
		readFile, err := os.ReadFile(file.AbsolutePath)
		if err != nil {
			panic(err)
		}
		s := string(readFile)
		if strings.Contains(s, "assets/087AC4D233B64EB0index.10284822.js") {
			fmt.Println(file.AbsolutePath)
			fileList = append(fileList, file.AbsolutePath)
		}
	}
}
