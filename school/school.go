package school

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Config struct {
	AppKey    string
	AppSecret string
	Ocode     string
}

type School struct {
	conf *Config
	AccessTokenHandle
}

func NewSchool(cfg *Config) *School {
	defaultAkHandle := NewAccessTokenHandle(cfg)
	return &School{conf: cfg, AccessTokenHandle: defaultAkHandle}
}

func GetConf() *Config {
	c := new(Config)
	yamlFile, err := ioutil.ReadFile("application.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}