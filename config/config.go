package config

import (
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v3"
)

type Conf struct {
	Dsn       string
	JwtSecret string `yaml:"jwt_secret"`
	Host      string
	Port      string
}

var Config *Conf

func InitConfig() error {
	yamlFile, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		return err
	}

	Config = &Conf{}

	err = yaml.Unmarshal(yamlFile, Config)

	spew.Dump("init conf", Config)

	return nil
}
