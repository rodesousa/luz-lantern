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

// YamlMapper : Implementation of Mapper interface for Yamlfiles
package mapper

import (
	"github.com/rodesousa/lantern/logger"
	"github.com/rodesousa/lantern/shard"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type ShardsAsMap []map[string]shard.ShardArguments
type ShardsAsYaml map[string]ShardsAsMap

//convert yaml to ShardsAsYaml Object
func MappingYaml(filename string) (ShardsAsYaml, error) {
	data, er := ioutil.ReadFile(filename)
	if er != nil {
		logger.FatalWithFields("Cannot read the file",
			logger.Fields{"filename": filename, "errors": er})
		return nil, er
	}

	shardsAsYaml := make(ShardsAsYaml)
	err := yaml.Unmarshal([]byte(data), &shardsAsYaml)

	if err != nil {
		logger.FatalWithFields("Error unmarshalling yaml file",
			logger.Fields{"filename": filename, "errors": er})
		return nil, err
	}
	return shardsAsYaml, nil
}

func AnalyseShard(in []map[string]shard.ShardArguments) []shard.Shard {
	var shards []shard.Shard
	// At this level, we are in the sub - cmd hierarchy
	for i := range in {
		for k, v := range in[i] {
			shards = append(shards, PatternMatching(k, v))
		}
	}
	return shards
}
