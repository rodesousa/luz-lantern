// Package shard provides ...
package shard

import (
	"github.com/rodesousa/lantern/logger"
	"github.com/rodesousa/lantern/utils"
	"os/exec"
	"strings"
)

func (cmd Ping) Cmd() bool {
	var toReturn = false
	cmdStatus, cmdMsg, error := exe_cmd(cmd.Cmd_line, cmd.Args["name"].(string))
	expected := utils.GetBool(cmd.Args["expected"], true)
	// If command == ok. Test cmdMsg
	if cmdStatus == true {
		toReturn = (strings.Contains(cmdMsg, cmd.Args["name"].(string)) == expected)
	} else {
		// Else test if result == expected in yaml conf file
		toReturn = (cmdStatus == expected)
	}
	logger.PrintShardResult("Shard Ping test result", cmdStatus, cmd.Args["name"].(string), cmdMsg, error)
	return toReturn
}

func exe_cmd(cmd []string, arg string) (bool, string, error) {
	var cmdTocall string
	var args string
	// build the command
	if len(cmd) == 1 {
		cmdTocall = cmd[0]
	} else {
		cmdTocall = cmd[0]
		args = strings.Join(cmd[1:len(cmd)], "")
	}
	// Launch the command
	var out []byte
	var err error
	// One arg cmd
	if args == "" {
		out, err = exec.Command(cmdTocall, arg).Output()
	} else {
		// > One args cmd
		out, err = exec.Command(cmdTocall, args, arg).Output()
	}
	if err != nil {
		logger.DebugWithFields("Error occured while testing command", logger.Fields{"cmd": cmd, "str_arg": arg, "str_error": err})
		return false, "", err
	}
	logger.DebugWithFields("Command ok", logger.Fields{"cmd": cmd, "str_arg": arg, "str_out": utils.ByteToString(out)})
	return true, utils.ByteToString(out), err
}
