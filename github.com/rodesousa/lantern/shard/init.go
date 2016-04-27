// age main provides ...
package shard

import (
	"container/list"
	"fmt"
	"github.com/rodesousa/lantern/logger"
	"os/exec"
	"runtime"
	"strings"
)

func InitUser() User {
	if runtime.GOOS == "windows" {
		return User{Shard{"user", "net user", make(Arg_type), list.New()}}
	} else {
		return User{Shard{"user", "id", make(Arg_type), list.New()}}
	}
}

func (cmd Shard) Cmd() bool {
	fmt.Println("Shard ! ")
	return true
}

func (cmd User) Cmd() bool {
	logger.Debug("Testing user")
	b := true
	//
	// list.List
	//
	if cmd.ArgsL.Len() != 0 {
		for e := cmd.ArgsL.Front(); e != nil; e = e.Next() {
			if !exe_cmd(cmd.Cmd_line, e.Value.(string)) {
				b = false
			}
		}
	}
	//
	// en shard.Arg_type
	//
	out, err := exec.Command(cmd.Cmd_line, cmd.Args["name"].(string)).Output()

	if err != nil {
		logger.Error("Error occured while testing command", logger.Fields{"cmd": cmd.Cmd_line, "str_arg": cmd.Args["name"].(string)})
		return false
	}
	logger.InfoWithFields("Command ok", logger.Fields{"cmd": cmd.Cmd_line, "str_arg": cmd.Args["name"].(string), "str_out": logger.ByteToString(out)})

	return b
}

func exe_cmd(cmd string, arg string) bool {
	parts := strings.Fields(cmd)
	size := len(parts)
	var cmdTocall string
	var args string
	// build the command
	if size == 1 {
		cmdTocall = cmd
	} else {
		cmdTocall = parts[0]
		for i := 1; i < len(parts); i += 1 {
			args = args + parts[i]
		}
	}

	out, err := exec.Command(cmdTocall, args, arg).Output()
	if err != nil {
		logger.Error("Error occured while testing command", logger.Fields{"cmd": cmd, "str_arg": arg})
		return false
	}
	logger.InfoWithFields("Command ok", logger.Fields{"cmd": cmd, "str_arg": arg, "str_out": logger.ByteToString(out)})

	return true
}
