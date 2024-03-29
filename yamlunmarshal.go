package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"log"
)



type Users struct {
	Name   string      `yaml:"name"`
	Age    int8        `yaml:"age"`
	Address  string `yaml:"address"`
	Hobby []string `yaml:"hobby"`
}


func main() {

	file, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var data [7]Users
	err2 := yaml.Unmarshal(file, &data)

	if err2 != nil {
		log.Fatal(err2)
	}
	for _, v := range data {
		fmt.Println(v)
	}
}
