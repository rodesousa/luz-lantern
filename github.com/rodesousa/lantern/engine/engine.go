// Package engine provides ...
package engine

import (
	"fmt"
	"github.com/rodesousa/lantern/logger"
	"github.com/rodesousa/lantern/shard"
	"gopkg.in/yaml.v2"
	"io/ioutil"
//"reflect"
	"container/list"
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
		logger.Fatal("Cannot read the file", logger.Fields{"errors": er})
	}
	// the Map for yaml unmarshalling
	mapYaml := make(map[string][]map[string]shard.Arg_type)
	// unmarshall
	err := yaml.Unmarshal([]byte(data), &mapYaml)
	// In case of error
	if err != nil {
		logger.Fatal("Error unmarshalling yaml file", logger.Fields{"errors": err})
	}
	// Launch the analysis
	for k, v := range mapYaml {
		// if it's a cmd, analyse the shards
		if k == "cmd" {
			analyseShard(v)
		}
	}
}

func analyseShard(in []map[string]shard.Arg_type) {
	// Shard List
	shards := list.New()
	// At this level, we are in the sub - cmd hierarchy
	for i := range in {
		for k, v := range in[i] {
			if k == "user" {
				// Case user
				shardUser := shard.InitUser()
				shardUser.Args = v
				shards.PushBack(shardUser)
			}  else if k == "ping" {
				//case ping
				shardPing := shard.InitPing()
				shardPing.Args = v
				shards.PushBack(shardPing)
			}
		}
	}

	// Launch test on shards
	for aShard := shards.Front(); aShard != nil; aShard = aShard.Next() {
		switch v := aShard.Value.(type) {
		case shard.User:
			v.Cmd()
		case shard.Ping:
			v.Cmd()
		}

	}
}
