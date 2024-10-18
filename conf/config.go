package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

var Config = new(AppConfig)

func Setup(reader ConfigReader) {
	config, err := reader.ReadConfig()
	if err != nil {
		panic(err)
	}
	s := string(config)
	fmt.Println(s)
	err = yaml.Unmarshal(config, &Config)
	if err != nil {
		panic(err)
	}
}

type Nats struct {
	Url string `json:"url" yaml:"url"`
}
type AppConfig struct {
	Nats Nats `json:"nats" yaml:"nats"`
}
