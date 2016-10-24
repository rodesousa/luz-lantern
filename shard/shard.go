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

type Check struct {
	Enabled bool
}

var CheckDisabled Check = Check{false}
var CheckEnabled Check = Check{true}

const (
	ValueChecked string = "Luz True"
)

type ShardArguments map[string]interface{}

type Result struct {
	Check bool
	Err   string
}

type Shard struct {
	Name             string
	Command          string
	CommandArguments []string
	Args             ShardArguments
	Status           Result
	Checked          Check
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
