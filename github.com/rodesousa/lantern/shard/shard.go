package shard

type Shard struct {
	name     string
	cmd_line string
}

type Cmd interface {
	Cmd() bool
}

type User struct {
	Shard
}
