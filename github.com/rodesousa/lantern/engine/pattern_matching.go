// Package engine provides ...
package engine

import (
	"github.com/rodesousa/lantern/shard"
)

func patternMatching(key string, value shard.Arg_type) shard.Shard {
	item := initShard(key)
	item.Args = value
	return item
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
