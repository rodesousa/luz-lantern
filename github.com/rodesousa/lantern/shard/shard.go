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
package shard

import (
	"runtime"
)

type ShardArguments map[string]interface{}

type Result struct {
	Check bool
	Err   error
}

type Shard struct {
	Name     string
	Cmd_line []string
	Args     ShardArguments
	Status   Result
}

type Cmd interface {
	Cmd() Result
}

func KoShards(shards []Shard) []Shard {
	new := make([]Shard, 0)
	for i := range shards {
		if shards[i].Status.Check == false {
			new = append(new, shards[i])
		}
	}
	return new
}

//
// INIT
//

var ResultDefault = Result{true, nil}

// USER
func InitUser() Shard {
	if runtime.GOOS == "windows" {
		return Shard{"user", []string{"net", "user"}, make(ShardArguments), ResultDefault}
	} else {
		return Shard{"user", []string{"id"}, make(ShardArguments), ResultDefault}
	}
}

// PING
func InitPing() Shard {
	return Shard{"ping", []string{"nslookup"}, make(ShardArguments), ResultDefault}
}

// UNKNOW
func InitUnknow() Shard {
	return Shard{"Unknow", []string{"???"}, make(ShardArguments), ResultDefault}
}
