// age main provides ...
package shard

import (
	"runtime"
)

func InitUser() Shard {
	if runtime.GOOS == "windows" {
		return Shard{"user", []string{"net", "user"}, make(Arg_type)}
	} else {
		return Shard{"user", []string{"id"}, make(Arg_type)}
	}
}

func InitPing() Shard {
	return Shard{"ping", []string{"nslookup"}, make(Arg_type)}
}

func InitUnknow() Shard {
	return Shard{"Unknow", []string{"???"}, make(Arg_type)}
}
