package webx

import (
	"fmt"
	"testing"
)

func TestReplaceFilePath(t *testing.T) {
	path, err := ReplaceFilePath("D:\\code\\git\\vite-test\\dist", "http://cdn.geeleo.com/web/beiluo/json")
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
}
