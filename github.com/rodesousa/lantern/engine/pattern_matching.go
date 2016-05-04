// Package engine provides ...
package engine

import (
	"container/list"
	"github.com/rodesousa/lantern/shard"
)

func patternMatching(key string, value shard.Arg_type, list *list.List) {
	item := initShard(key)
	item.Args = value
	list.PushBack(item)
}

func initShard(key string) shard.Shard {
	switch key {
	case "user":
		return shard.InitUser()
	case "ping":
		return shard.InitPing()
	}
	return shard.InitUnknow()
}
