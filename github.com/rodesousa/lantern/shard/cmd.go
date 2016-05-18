// Package shard provides ...
package shard

import (
	"os/exec"
	"strings"
)

func getExpected(cmd *Shard) bool {
	if val, ok := cmd.Args["expected"]; ok {
		return val.(bool)
	}
	return true
}

func (cmd *Shard) Cmd() bool {
	status, error := exeCmd(cmd.Cmd_line, cmd.Args["name"].(string))
	if getExpected(cmd) != status {
		cmd.Err = error
		cmd.status = false
	}
	return cmd.status
}

func exeCmd(cmd []string, arg string) (bool, error) {
	var cmdTocall, args string
	var out []byte
	var err error

	// build the command
	cmdTocall = cmd[0]
	if len(cmd) != 1 {
		args = strings.Join(cmd[1:len(cmd)], "")
	}

	// One args or more
	if args == "" {
		out, err = exec.Command(cmdTocall, arg).Output()
	} else {
		out, err = exec.Command(cmdTocall, args, arg).Output()
	}

	if out != nil { //TODO
	}

	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
