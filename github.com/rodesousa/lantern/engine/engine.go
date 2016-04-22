// Package engine provides ...
package engine

import (
	"fmt"
	"github.com/rodesousa/lantern/shard"
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
	mapYaml := make(map[string][]map[string]map[interface{}]interface{})

	if er != nil {
		fmt.Println("Cannot read the file")
		os.Exit(1)
	}

	err := yaml.Unmarshal([]byte(data), &mapYaml)

	if err != nil {
		fmt.Println("error: %v", err)
	}

	value := mapYaml["cmd"][0]

	for k, v := range value {
		if k == "user" {
			shard := shard.InitUser()
			for k2, v2 := range v {
				shard.Args[k2.(string)] = []string{v2.(string)}
			}
		}
	}
}
