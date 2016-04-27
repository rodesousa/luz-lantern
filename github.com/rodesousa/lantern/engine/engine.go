// Package engine provides ...
package engine

import (
	"fmt"
	"github.com/rodesousa/lantern/shard"
	"gopkg.in/yaml.v2"
	"io/ioutil"
//"reflect"
	"github.com/rodesousa/lantern/logger"
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
	// read the file from commande line
	data, er := ioutil.ReadFile(filename)
	// In case of error
	if er != nil {
		logger.Fatal("Cannot read the file", logger.Fields{"errors" : er})
	}
	// the Map for yaml unmarshalling
	mapYaml := make(map[string]interface{})
	// unmarshall
	err := yaml.Unmarshal([]byte(data), &mapYaml)
	// In case of error
	if err != nil {
		logger.Fatal("Error unmarshalling yaml file", logger.Fields{"errors" : err})
	}
	// Launch the analysis
	for k, v := range mapYaml {
		// if it's a cmd, analyse the shards
		if k == "cmd" {
			analyseShard(v)
		}
	}
}

func analyseShard(in interface{}) {
	local := in.([]interface{})
	shardUser := shard.InitUser()
	// At this level, we are in the sub - cmd hierarchy
	for i := range local {
		// convert the subShard for analysis
		subShard := local[i].(map[interface{}]interface{})
		// anayse the shard
		for k, v := range subShard {
			// Case user
			if k == "user" {
				name, error := extractString(v, "name")
				if (error == nil) {
					shardUser.ArgsL.PushBack(name)
				} else {
					logger.Error("Error extractString", logger.Fields{"errors" : error})
					//fmt.Println(error.Error())
				}
			}
		}
	}
	//fmt.Println(shardUser)
	// Launch test on users
	shardUser.Cmd()
}

func extractString(in interface{}, key string) (string, error) {
	switch v := in.(type) {
	case map[interface{}]interface{} :
		for k, v := range v {
			if (k == key) {
				return v.(string), nil
			}
		}
		return "", fmt.Errorf("Unable to find %#v in %#v", key, in)
	default:
		return "", fmt.Errorf("Unable to Cast %#v to string", in)
	}
}
