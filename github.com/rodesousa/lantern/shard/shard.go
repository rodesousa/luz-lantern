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
	"errors"
	"fmt"
	"runtime"
)

type ShardArguments map[string]interface{}

type Result struct {
	Check bool
	Err   string
}

type Shard struct {
	Name             string
	Command          string
	CommandArguments string
	Args             ShardArguments
	Status           Result
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

type MapString map[string]string

func (m MapString) exists(find string) bool {
	_, ok := m[find]
	return ok
}

func newErrorArg(err string) error {
	return errors.New(fmt.Sprintf("Value %s doesnt exist", err))
}

func (m ShardArguments) nameExist(name string) string {
	if v, ok := m["name"]; ok {
		return fmt.Sprintf("%s %s", name, v)
	} else {
		return name
	}
}

func (m ShardArguments) argsExist(name string) error {
	if _, ok := m[name]; ok {
		return nil
	} else {
		return newErrorArg(name)
	}
}

var ResultDefault = Result{true, ""}

// USER
func InitUser(args ShardArguments) (error, Shard) {

	if runtime.GOOS == "windows" {
		//return Shard{"user", []string{"net", "user"}, value, ResultDefault}
		return errors.New("not implem"), Shard{}
	} else {
		name := args.nameExist("user")

		var cmd, cmdArgs string
		if err := args.argsExist("name"); err == nil {
			cmd = "id"
			cmdArgs = args["name"].(string)
		} else {

			return err, Shard{}
		}

		return nil, Shard{name, cmd, cmdArgs, args, ResultDefault}
	}
}

// PING
func InitPing(args ShardArguments) (error, Shard) {
	name := args.nameExist("ping")

	var cmd, cmdArgs string
	if err := args.argsExist("url"); err == nil {
		cmd = "nslookup"
		cmdArgs = args["url"].(string)
	} else {
		return err, Shard{}
	}

	return nil, Shard{name, cmd, cmdArgs, args, ResultDefault}
}

// UNKNOW
func InitUnknow() (error, Shard) {
	return errors.New("not implem"), Shard{}
	//return Shard{"Unknow", []string{"???"}, make(ShardArguments), ResultDefault}
}
