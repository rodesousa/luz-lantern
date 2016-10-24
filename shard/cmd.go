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
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func getExpected(shard *Shard) bool {
	if val, ok := shard.Args["expected"]; ok {
		return val.(bool)
	}
	return true
}

func (s *Shard) Cmd() bool {
	out, status, err := s.exeCmd()
	if getExpected(s) != status {
		if err != "" {
			s.Status.Err = err
			s.Status.Check = false
		}

		if s.Checked.Enabled {
			if strings.Contains(string(out[:]), ValueChecked) != false {
				s.Status.Err = fmt.Sprintf("\n %s", out)
				s.Status.Check = false
			}
		}
	}
	return s.Status.Check
}

func (s Shard) exeCmd() (string, bool, string) {
	c := exec.Command(s.Command, s.CommandArguments...)

	var stderr bytes.Buffer
	c.Stderr = &stderr

	err := c.Run()

	stdout := fmt.Sprint(os.Stdout)

	if err != nil {
		return stdout, false, stderr.String()
	} else {
		return stdout, true, ""
	}
}
