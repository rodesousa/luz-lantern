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

package controller

import (
	"fmt"
	"github.com/rodesousa/lantern/engine"
	log "github.com/rodesousa/lantern/logger"
	"github.com/rodesousa/lantern/mapper"
	"github.com/rodesousa/lantern/shard"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

//Controller
type Controller struct {
	filename string
}

var controller Controller

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:     "run [yaml_file]",
	Short:   "launch the lantern program",
	Long:    `launch the lanter program with a yaml file that discribe all test to do`,
	Example: "lantern run tests.yaml",
	Run:     runLuz,
}

// Main funtion, launch when run command is invoked
func runLuz(cmd *cobra.Command, args []string) {
	// waiting for 1 arg -> show help in this case
	if len(args) == 0 {
		cmd.SetOutput(log.GetOutLogger())
		cmd.Help()
		os.Exit(1)
	} else {
		log.Info("Starting lantern with run command")
		log.DebugWithFields("run called", log.Fields{"args": args})
		controller = Controller{args[0]}
		// Launch lantern in selected mode
		launchLantern()

	}
	log.Info("End lantern")
}

// Main method of the lantern program
func launchLantern() {
	// Init the mapper
	shardsAsYaml, err := mapper.MappingYaml(controller.filename)
	if err == nil {
		// Get the shards from the Mapper
		shards := mapper.AnalyseShard(shardsAsYaml["cmd"])
		// Call the Engine with the shards
		engine.RunMultiThread(shards)

		koShards := shard.KoShards(shards)
		ko := len(koShards)
		ok := len(shards) - len(koShards)
		sOk := strconv.Itoa(ok) + "/" + strconv.Itoa(len(shards))
		sKo := strconv.Itoa(ko) + "/" + strconv.Itoa(len(shards))
		log.InfoWithFields("Test OK", log.Fields{"nbOk": sOk})
		log.InfoWithFields("Test KO", log.Fields{"nbKO": sKo})

		if ko > 0 {
			for i := range koShards {
				shard := koShards[i]
				err := fmt.Sprintf("%s : %s", shard.Name, shard.Status.Err)
				log.Info(err)
			}
		}
	} else {
		//error in mapping yaml
		//TODO
	}
}

func init() {
	RootCmd.AddCommand(runCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
