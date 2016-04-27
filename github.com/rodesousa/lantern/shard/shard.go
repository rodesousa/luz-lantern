package shard

import "container/list"

type Arg_type map[string]interface{}

type Shard struct {
	Name     string
	Cmd_line string
	Args     Arg_type
	ArgsL    *list.List
}

type Cmd interface {
	Cmd() bool
}

type User struct {
	Shard
}
