// Package shard provides ...
package shard

import (
	"fmt"
	"github.com/rodesousa/lantern/logger"
	"github.com/rodesousa/lantern/utils"
	"os/exec"
	"strings"
)

func getExpected(cmd Shard) bool {
	if val, ok := cmd.Args["expected"]; ok {
		return val.(bool)
	}
	return false
}

// todo
// a mettre ds logger
func printCmdResult(expected bool, cmdStatus bool, cmd Shard) string {
	if expected == cmdStatus {
		if cmdStatus == false {
			return "Test Ok: " + cmd.Name + " " + cmd.Args["name"].(string)
		}
		return "Test KO but expected KO: " + cmd.Name + " " + cmd.Args["name"].(string)
	}
	if cmdStatus == false {
		return "Test KO: " + cmd.Name + " " + cmd.Args["name"].(string)
	}
	return "Test Ok but expected KO : " + cmd.Name + " " + cmd.Args["name"].(string)
}

func (cmd Shard) Cmd() bool {
	cmdStatus, cmdMsg, error := exeCmd(cmd.Cmd_line, cmd.Args["name"].(string))
	//look expected argument
	expected := getExpected(cmd)
	msg := printCmdResult(expected, cmdStatus, cmd)
	fmt.Println(msg)

	if (error != nil) && (cmdMsg != "") {
		// pour eviter le unsed... a retirer !
	}

	return expected == cmdStatus
}

func exeCmd(cmd []string, arg string) (bool, string, error) {
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
