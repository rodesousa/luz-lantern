package main

import (
	"fmt"
	//"github.com/rodesousa/lantern/controller"
	//"github.com/rodesousa/lantern/shard"
	"io/ioutil"
	"os"
)

func main() {
	file := os.Args[1:]
	dat, err := ioutil.ReadFile(file[0])

	if err != nil {
		fmt.Println("Cannot read the file")
		os.Exit(1)
	}

	//var u = shard.InitUser()
	//fmt.Println(controller.Test())
	//fmt.Println(u.Cmd())

	fmt.Println(string(dat))
}
