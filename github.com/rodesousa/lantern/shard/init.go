// age main provides ...
package shard

import (
	"github.com/rodesousa/lantern/logger"
	"os/exec"
	"runtime"
	"strings"
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
	cmdStatus, _ := exe_cmd(cmd.Cmd_line, cmd.Args["name"].(string))
	logger.DebugWithFields("Return status of command", logger.Fields{"return" : cmdStatus})
	return cmdStatus
}

func (cmd Ping) Cmd() bool {
	var toReturn = false;
	cmdStatus, cmdMsg := exe_cmd(cmd.Cmd_line, cmd.Args["name"].(string))
	expected := getBool(cmd.Args["expected"], true)
	// If command == ok. Test cmdMsg
	if cmdStatus == true {
		toReturn = (strings.Contains(cmdMsg, cmd.Args["name"].(string)) == expected)
	} else {
		// Else test if result == expected in yaml conf file
		toReturn = (cmdStatus == expected)
	}
	logger.DebugWithFields("Return status of command", logger.Fields{"return" : toReturn})
	return toReturn
}

func getBool(toAnalyse interface{}, defaultValue bool) bool {
	if toAnalyse == nil {
		return defaultValue
	}
	return toAnalyse.(bool)
}

func exe_cmd(cmd []string, arg string) (bool, string) {
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
		logger.Error("Error occured while testing command", logger.Fields{"cmd": cmd, "str_arg": arg, "str_error" : err})
		return false, err.Error()
	}
	logger.InfoWithFields("Command ok", logger.Fields{"cmd": cmd, "str_arg": arg, "str_out": logger.ByteToString(out)})

	return true, logger.ByteToString(out)
}
