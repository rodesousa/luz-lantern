// Package main provides ...
package shard

import (
	"fmt"
	"container/list"
	"runtime"
	"os/exec"
	"strings"
	"github.com/rodesousa/lantern/logger"
)

func InitUser() User {
	if runtime.GOOS == "windows" {
		return User{Shard{"user", "net user", make(map[string][]string), list.New()} }
	} else {
		return User{Shard{"user", "id", make(map[string][]string), list.New()} }
	}
}

func (cmd Shard) Cmd() bool {
	fmt.Println("Shard ! ")
	return true
}

func (cmd User) Cmd() bool {
	logger.Debug("Testing user")
	b := true
	if (cmd.ArgsL.Len() != 0) {
		for e := cmd.ArgsL.Front(); e != nil; e = e.Next() {
			if ! exe_cmd(cmd.Cmd_line, e.Value.(string)) {
				b = false
			}
		}
	}
	return b
}

func exe_cmd(cmd string, arg string) bool {
	//fmt.Println(cmd)
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
		logger.Error("Error occured while testing command", logger.Fields{"cmd" : cmd, "str_arg" : arg})
		return false
	}
	logger.InfoWithFields("Command ok", logger.Fields{"cmd" : cmd, "str_arg" : arg, "str_out" : logger.ByteToString(out)})

	return true
}