package os

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Service struct {
	AppMode   string `yaml:"AppMode"`
	HttpPort  string `yaml:"HttpPort"`
}

type MySql struct {
	Db 			string `yaml:"Db"`
	DbHost		string `yaml:"DbHost"`
	DbPort		string `yaml:"DbPort"`
	DbPassword  string `yaml:"DbPassword"`
	DbName		string `yaml:"DbName"`
	DbUser		string `yaml:"DbUser"`
	Create 		int 	`yaml:"Create"`
	CreatePath  string  `yaml:"CreatePath"`
}

type Redis struct {
	Address  		string  `yaml:"Address"`
	Password    string `yaml:"Password"`
}

type Web3 struct {
	Address string `yaml:"Address"`
	HttpUrl string `yaml:"HttpUrl"`
}

type WebType struct {
	Debug Web3 `yaml:"debug"`
}

type Yaml struct {
	Name     string `yaml:"name"`
	Service  Service `yaml:"service"`
	MySql    MySql `yaml:"mysql"`
	Redis	 Redis `yaml:"redis"`
	Web3     WebType  `yaml:"web3"`
}
var YamlResult *Yaml
func OsYaml()  {
	_path, err := os.Getwd()
	if err == nil {
		_path += "/pubspec.yaml"
	}
	yamlFile, err := ioutil.ReadFile(_path)

	if err == nil {
		fmt.Println()
		yaml.Unmarshal(yamlFile, &YamlResult)
		fmt.Println(YamlResult.MySql.Db)
	}else {
		fmt.Println(err, "错误")
	}
}
