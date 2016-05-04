// age main provides ...
package shard

import (
	"github.com/rodesousa/lantern/logger"
	"runtime"
)

func InitUser() User {
	if runtime.GOOS == "windows" {
		return User{Shard{"user", []string{"net", "user"}, make(Arg_type)}}
	} else {
		return User{Shard{"user", []string{"id"}, make(Arg_type)}}
	}
}

func InitPing() Ping {
	return Ping{Shard{"ping", []string{"nslookup"}, make(Arg_type)}}
}

func (cmd User) Cmd() bool {
	cmdStatus, cmdMsg, error := exe_cmd(cmd.Cmd_line, cmd.Args["name"].(string))
	logger.PrintShardResult("Shard User test result", cmdStatus, cmd.Args["name"].(string), cmdMsg, error)
	return cmdStatus
}
