package settings

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var configs Settings

type (
	Settings struct {
		Application
		Database
		Cache
		Rpc    Rpc
		Consul Consul
	}

	Application struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Ssl  int    `yaml:"ssl"`
	}

	Database struct {
		Source string `yaml:"source"`
	}

	Cache struct {
		Redis string `yaml:"redis"`
	}

	Rpc struct {
		Name string `yaml:"name"`
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	Consul struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
)

func init() {
	data, err := ioutil.ReadFile("./etc/config.dev.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &configs)
	if err != nil {
		panic(err)
	}
}

// GetSettings 获取基础配置
func GetSettings() Settings {
	return configs
}
