package shard

type Arg_type map[string]interface{}

type Shard struct {
	Name     string
	Cmd_line []string
	Args     Arg_type
}

type Cmd interface {
	Cmd() bool
}

type User struct {
	Shard
}

type Ping struct {
	Shard
}