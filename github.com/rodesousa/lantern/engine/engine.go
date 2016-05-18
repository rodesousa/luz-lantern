// Package engine provides ...
package engine

import (
	"fmt"
	"github.com/rodesousa/lantern/logger"
	"github.com/rodesousa/lantern/shard"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
	if er != nil {
		logger.FatalWithFields("Cannot read the file", logger.Fields{"errors": er})
	}
	mapYaml := make(map[string][]map[string]shard.Arg_type)
	err := yaml.Unmarshal([]byte(data), &mapYaml)
	if err != nil {
		logger.FatalWithFields("Error unmarshalling yaml file", logger.Fields{"errors": err})
	}
	// Launch the analysis
	for k, v := range mapYaml {
		if k == "cmd" {
			analyseShard(v)
		}
	}
}

func analyseShard(in []map[string]shard.Arg_type) {
	shards := make([]shard.Shard, len(in))
	// Built yaml to object
	for i := range in {
		for k, v := range in[i] {
			shards[i] = patternMatching(k, v)
		}
	}
	for i := range shards {
		p := &shards[i]
		p.Cmd()
	}

	ko := shard.KoShards(shards)
	ok := len(shards) - len(ko)
	fmt.Println("Test Ok :", ok, "/", len(shards))

	fmt.Println("Test KO :", len(ko), "/", len(shards))

	for i := range ko {
		shard := ko[i]
		fmt.Println(shard.Name+" : ", shard.Err)
	}
}
