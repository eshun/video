package models

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"flag"
	"fmt"
)

type Fvip struct {
	Name    string `yaml:"name"`
	Url     string `yaml:"url"`
	Invalid bool   `yaml:"invalid"`
}

type Conf struct {
	Port int    `yaml: "port"`
	Fvip []Fvip `yaml: "fvip"`
}

func GetConf() (*Conf, error) {
	confPath := flag.String("c", "./conf/conf.yaml","confPath")
	flag.Parse()

	fmt.Println(*confPath)

	file, err := ioutil.ReadFile(*confPath)
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
