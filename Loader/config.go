package Loader

import (
	"fmt"
	v2 "gopkg.in/yaml.v2"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
}

var conf *Config

func init() {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	conf = new(Config)
	if err := v2.Unmarshal(data, conf); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}

func (this *Config) String() string {
	yamlData, err := yaml.Marshal(conf)
	if err != nil {
		return ""
	}
	return string(yamlData)
}

func GetConfig() *Config {
	return conf
}

func SaveConfig() {
	data, err := yaml.Marshal(&conf)
	if err != nil {
		fmt.Printf("无法序列化配置：%v\n", err)
		return
	}
	if err := os.WriteFile("./config.yaml", data, 0644); err != nil {
		fmt.Printf("无法写入配置文件：%v\n", err)
		return
	}
	return
}
