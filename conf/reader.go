package conf

import (
	"os"
)

type ConfigReader interface {
	ReadConfig() ([]byte, error)
}

type LocalConfig struct {
	FilePath string
}

func (conf *LocalConfig) ReadConfig() ([]byte, error) {
	return os.ReadFile(conf.FilePath)
}
