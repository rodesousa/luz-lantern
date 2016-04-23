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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
	"github.com/rodesousa/lantern/engine"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [yaml_file]",
	Short: "launch the lantern program",
	Long: `launch the lanter program with a yaml file that discribe all test to do`,
	Example : "lantern run tests.yaml",

	Run: runLuz,

}

func runLuz(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println(cmd.Help())
		os.Exit(1)
	} else  {
		fmt.Println("run called with", args)
		engine.MapYamlToShard(args[0])

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
