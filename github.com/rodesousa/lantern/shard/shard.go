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

// deprecated
type User Shard

// deprecated
type Ping Shard
