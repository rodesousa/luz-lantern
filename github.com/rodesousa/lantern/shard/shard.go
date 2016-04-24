package shard

import "container/list"

type arg_type map[string][]string

type Shard struct {
	Name     string
	Cmd_line string
	Args     arg_type
	ArgsL    *list.List
}

type Cmd interface {
	Cmd() bool
}

type User struct {
	Shard
}
