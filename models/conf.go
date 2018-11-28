package models

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type conf struct {
	Port int `yaml: "port"`
}
var Conf = conf{
	Port: 8080,
}

func GetConf() (*conf, error) {
	file, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(file, &Conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return &Conf, err
}