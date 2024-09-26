package compress

import (
	"fmt"
	"github.com/mholt/archiver/v3"
	"os"
)

// DeCompress 使用 archiver 解压缩文件
func DeCompress(compressFilePath, dest string) error {
	// 创建输出目录
	if err := os.MkdirAll(dest, 0755); err != nil {
		return fmt.Errorf("DeCompress error creating output directory: %w", err)
	}

	// 使用 archiver 解压缩
	if err := archiver.Unarchive(compressFilePath, dest); err != nil {
		return fmt.Errorf("DeCompress error unarchiving file: %w", err)
	}

	return nil
}
