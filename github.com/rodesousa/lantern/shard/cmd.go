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
)

func getExpected(shard *Shard) bool {
	if val, ok := shard.Args["expected"]; ok {
		return val.(bool)
	}
	return true
}

func (s *Shard) Cmd() bool {
	status, error := s.exeCmd()
	if getExpected(s) != status {
		s.Status.Err = error
		s.Status.Check = false
	}
	return s.Status.Check
}

type FnEmpty func() string

func (shard *Shard) Do(fn FnEmpty) string {
	return fn()
}

func (s Shard) exeCmd() (bool, error) {
	var out []byte
	var err error

	// One args or more
	out, err = exec.Command(s.Command, s.CommandArguments).Output()

	if out != nil { //TODO
	}

	if err != nil {
		return false, err
	} else {
		return true, err
	}
}
