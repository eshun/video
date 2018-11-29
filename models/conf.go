package models

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

type Fvip struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
	Invalid bool `yaml:"invalid"`
}

type Conf struct {
	Port int  `yaml: "port"`
	Fvip []Fvip `yaml: "fvip"`
}

func GetConf() (*Conf, error) {
	dbPath, err := filepath.Abs("./conf/conf.yaml")
	if err != nil {
		log.Fatalf("yaml not find conf: %v", err)
	}
	file, err := ioutil.ReadFile(dbPath)
	if err != nil {
		log.Fatalf("yaml not read: %v", err)
	}
	c := Conf{
		Port: 8080,
	}

	err = yaml.Unmarshal(file, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return &c, err
}
