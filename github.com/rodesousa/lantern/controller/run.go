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
	"encoding/json"
	"fmt"
	"github.com/rodesousa/lantern/engine"
	log "github.com/rodesousa/lantern/logger"
	"github.com/rodesousa/lantern/mapper"
	"github.com/rodesousa/lantern/shard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"strconv"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "lantern",
	Short: "lantern is a tool for testing environments",
	Long: `===============================================================================
    __                         __                    __
   / /   __  __ ____          / /   ____ _   ____   / /_  ___    _____   ____
  / /   / / / //_  / ______  / /   / __  /  / __ \ / __/ / _ \  / ___/  / __ \
 / /___/ /_/ /  / /_/_____/ / /___/ /_/ /  / / / // /_  /  __/ / /     / / / /
/_____/\__,_/  /___/       /_____/\__,_/  /_/ /_/ \__/  \___/ /_/     /_/ /_/

===============================================================================

luz-lantern is a tool program used to test/check environments from development to production.

Please check down for help.

Copyright © 2016
Roberto De Sousa (https://github.com/rodesousa)
Patrick Tavares (https://github.com/ptavares)
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

}

//Controller
type Controller struct {
	filename string
	result   string
}

var (
	cfgFile    string
	logFile    string
	debug      bool
	off        bool
	controller Controller
)

var runCmd = &cobra.Command{
	Use:     "run [yaml_file]",
	Short:   "launch the lantern program",
	Long:    `launch the lanter program with a yaml file that discribe all test to do`,
	Example: "lantern run tests.yaml",
	Run:     runLuz,
}

var serverCmd = &cobra.Command{
	Use:     "server [yaml_file]",
	Short:   "launch the lantern program like a server",
	Long:    `launch the lanter program with a yaml file that discribe all test to do`,
	Example: "lantern server tests.yaml",
	Run:     lanternServer,
}

func init() {
	cobra.OnInitialize(initFromCL)
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(serverCmd)

	RootCmd.PersistentFlags().StringVar(&logFile, "logfile", "", "log file output (default is current path)")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "show debug message")
	RootCmd.PersistentFlags().BoolVarP(&off, "off", "o", false, "disable out console log")
}

// initConfig reads in config file and ENV variables if set.
func initFromCL() {

	////mode server, disabled out console log
	//if server {
	//	off = true
	//}

	// initialize debug level
	log.Init(debug, !off, (logFile != ""), logFile)

	if cfgFile != "" {
		// enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".luz-lantern") // name of config file (without extension)
	viper.AddConfigPath("$HOME")        // adding home directory as first search path
	viper.AutomaticEnv()                // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		RootCmd.SetOutput(log.GetOutLogger())
		RootCmd.Help()
		os.Exit(-1)
	}
}

func lanternServer(cmd *cobra.Command, args []string) {
	runLuz(cmd, args)
	controller.runServer()
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

		// response json
		mapB, _ := json.Marshal(map[string]string{"status": "OK"})
		controller = Controller{args[0], string(mapB)}

		// Launch lantern in selected mode
		controller.launchLantern()
	}
	log.Info("End lantern")
}

// Main method of the lantern program
func (controller *Controller) launchLantern() {
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

		// lantern find check ko
		if ko > 0 {
			err := fmt.Sprintf("Test KO nbKO=%s\n", sKo)
			mapResult := map[string]string{}
			mapResult["status"] = "KO"
			for i := range koShards {
				shard := koShards[i]
				err = fmt.Sprintf("%s - %s : %s", err, shard.Name, shard.Status.Err)
				mapResult[shard.Name] = shard.Status.Err
				mapB, _ := json.Marshal(mapResult)
				controller.result = string(mapB)
			}
			log.Info(err)
		}

	} else {
		//error in mapping yaml
		//TODO
	}
}

func (controller Controller) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, controller.result)
}

func (controller Controller) runServer() {
	http.HandleFunc("/", controller.handler)
	http.ListenAndServe(":8080", nil)
}
