package configs

import (
	"fmt"
	"io/ioutil"

	"sync"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DEBUG             bool   `yaml:"debug"`
	DOCKER_REMOTE_URI string `yaml:"docker_remote_uri"`
	DATABASE_BACKEND  string `yaml:"database_backend"`
	DATABASE_URI      string `yaml:"database_uri"`
	REALM             string `yaml:"realm"`
	SCRETKEY          string `yaml:"scret_key"`
	PORT              string `yaml:"port"`
}

var (
	conf *Config
	once sync.Once
)

func Conf() *Config {
	once.Do(func() {
		conf = &Config{}
		if confData, err := ioutil.ReadFile("./configs/conf.yaml"); err == nil {
			yaml.Unmarshal(confData, conf)
		}
	})
	fmt.Println(conf)
	return conf
}
