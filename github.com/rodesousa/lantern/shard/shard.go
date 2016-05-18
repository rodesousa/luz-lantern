package shard

type Arg_type map[string]interface{}

type Shard struct {
	Name     string
	Cmd_line []string
	Args     Arg_type
	Err      error
	status   bool
}

type Cmd interface {
	Cmd() bool
}

func KoShards(shards []Shard) []Shard {
	new := make([]Shard, 0)
	for i := range shards {
		if shards[i].status == false {
			new = append(new, shards[i])
		}
	}
	return new
}

// deprecated
type User Shard

// deprecated
type Ping Shard
