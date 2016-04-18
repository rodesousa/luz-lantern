package main

import (
	"fmt"
	"github.com/rodesousa/lantern/controller"
	"github.com/rodesousa/lantern/shard"
)

func main() {
	var u = shard.InitUser()
	fmt.Println(controller.Test())
	fmt.Println(u.Cmd())
}
