package main

import (
	"fmt"
	//"github.com/rodesousa/lantern/controller"
	//"github.com/rodesousa/lantern/shard"
	"github.com/rodesousa/lantern/engine"
	//	"gopkg.in/yaml.v2"
	//"io/ioutil"
	//"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Il faut que 2 arguments !")
		os.Exit(1)
	}
	engine.MapYamlToShard(os.Args[1])
}
