package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	MongoURL string `json:"mongoURL"`
	MongoDB  string `json:"mongoDB"`
	Port     int    `json:"port"`
}

func NewConfig(file string) (*Config, error) {
	_, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = json.Unmarshal(f, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
