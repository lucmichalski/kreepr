package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type conf struct {
	DbUser		 string `yaml:"user"`
	DbPassword	 string `yaml:"password"`
	DbSchema	 string `yaml:"schema"`
	Key          string `yaml:"key"`
	Secret       string `yaml:"secret"`
	Token        string `yaml:"token"`
	AccessToken  string `yaml:"access_token"`
	AccessSecret string `yaml:"access_secret"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("../secrets.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
