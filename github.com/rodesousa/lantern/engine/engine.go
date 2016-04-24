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
	// read the file from commande line
	data, er := ioutil.ReadFile(filename)
	// In case of error
	if er != nil {
		fmt.Println("Cannot read the file")
		os.Exit(1)
	}
	// the Map for yaml unmarshalling
	mapYaml :=  make (map[string]interface{})
	// unmarshall
	err := yaml.Unmarshal([]byte(data), &mapYaml)
	// In case of error
	if err != nil {
		fmt.Println("error: %v", err)
		os.Exit(1)
	}
	// Launch the analysis
	for k,v := range mapYaml {
		// if it's a cmd, analyse the shards
		if k == "cmd" {
			analyseShard(v)
		}
	}
}

func analyseShard(in interface{}){
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
				if(error == nil) {
					//fmt.Println("Find a user Shard, extracted name : ", name)
					shardUser.ArgsL.PushBack(name)
				} else {
					fmt.Println(error.Error())
				}
			}
		}
	}
	//fmt.Println(shardUser)
	// Launch test on users
	shardUser.Cmd()
}

func extractString(in interface {}, key string) (string, error){
	switch v := in.(type) {
	case map[interface {}]interface {} :
		for k,v := range v {
			if (k == key) {
				return v.(string), nil
			}
		}
		return "", fmt.Errorf("Unable to find %#v in %#v", key, in)
	default:
		return "", fmt.Errorf("Unable to Cast %#v to string", in)
	}
}

func MapYamlToShard2(filename string) {
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
		fmt.Println("-----> key "+k)
		if k == "user" {
			shard := shard.InitUser()
			for k2, v2 := range v {
				shard.Args[k2.(string)] = []string{v2.(string)}
				fmt.Println(shard.Args)

			}
		}
	}
}
