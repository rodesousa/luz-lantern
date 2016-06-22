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

// Package engine provides the different engine possible
package engine

import (
	"github.com/rodesousa/lantern/shard"
	"runtime"
	"sync"
)

func RunMultiThread(shards []shard.Shard) []shard.Shard {
	// Number of available cores
	numCore := runtime.NumCPU()
	c := make(chan bool, numCore-1)
	wg := new(sync.WaitGroup)
	// Number of goroutines
	wg.Add(len(shards))
	// call the goroutines
	for i := range shards {
		go callShardExec(c, wg, &shards[i])
		// blocking c at each iteration
		c <- true
	}
	// Wait for all the children to die
	wg.Wait()
	close(c)
	return shards
}

func callShardExec(c chan bool, wg *sync.WaitGroup, shard *shard.Shard) {
	defer func() {
		<-c
		wg.Done() // Decrease the number of alive goroutines
	}()
	shard.Cmd()
}
