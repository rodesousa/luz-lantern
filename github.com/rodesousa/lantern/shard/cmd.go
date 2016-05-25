// Copyright © 2016 Roberto De Sousa (https://github.com/rodesousa) / Patrick Tavares (https://github.com/ptavares)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
		cmd.Status.Err = error
		cmd.Status.Check = false
	}
	return cmd.Status.Check
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
