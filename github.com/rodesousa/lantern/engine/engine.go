// Package engine provides ...
package engine

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	//"reflect"
)

type Engine struct {
}

type Run interface {
	Run() bool
}

func (engine Engine) Run() bool {
	fmt.Println("Engine ! ")
	return true
}

func MapYamlToShard(filename string) {
	data, er := ioutil.ReadFile(filename)
	mapYaml := make(map[string][]map[string]interface{})

	if er != nil {
		fmt.Println("Cannot read the file")
		os.Exit(1)
	}

	err := yaml.Unmarshal([]byte(data), &mapYaml)

	if err != nil {
		fmt.Println("error: %v", err)
	}
}
