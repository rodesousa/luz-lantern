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
	"github.com/rodesousa/luz-lantern/engine"
	log "github.com/rodesousa/luz-lantern/logger"
	"github.com/rodesousa/luz-lantern/mapper"
	"github.com/rodesousa/luz-lantern/shard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

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
}

// // //
//
// Init Arg
//
// // //

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
	Use:     "run",
	Short:   "launch the lantern program",
	Long:    `launch the lanter program with a yaml file that discribe all test to do`,
	Example: "lantern run",
	Run:     lantern,
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "launch the lantern program like a server",
	Long:  `launch the lanter program with a yaml file that discribe all test to do`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Miss a argument, lantern server status | start | stop")
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "launch the lantern program like a server",
	Long:  `launch the lanter program with a yaml file that discribe all test to do`,
	Run:   serverStart,
}

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "launch the lantern program like a server",
	Long:    `launch the lanter program with a yaml file that discribe all test to do`,
	Example: "lantern server tests.yaml",
	Run:     serverStatus,
}

var stopCmd = &cobra.Command{
	Use:     "stop",
	Short:   "launch the lantern program like a server",
	Long:    `launch the lanter program with a yaml file that discribe all test to do`,
	Example: "lantern server tests.yaml",
	Run:     serverStop,
}

func init() {
	cobra.OnInitialize(initFromCL)
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(serverCmd)
	serverCmd.AddCommand(stopCmd)
	serverCmd.AddCommand(statusCmd)
	serverCmd.AddCommand(startCmd)

	RootCmd.PersistentFlags().StringVar(&logFile, "logfile", "", "log file output (default is current path)")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "show debug message")
	RootCmd.PersistentFlags().BoolVarP(&off, "off", "o", false, "disable out console log")
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "conf.yaml", "conf file")
}

func initFromCL() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigName(".luz-lantern")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	if debug {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		RootCmd.SetOutput(log.GetOutLogger())
		RootCmd.Help()
		os.Exit(-1)
	}
}

// // //
//
// Main Program
//
// // //

func serverStart(cmd *cobra.Command, args []string) {
	log.Init(debug, false, (logFile != ""), logFile)
	runLuz(cmd, args)
	controller.runServer()
}

func serverStatus(cmd *cobra.Command, args []string) {
	log.Init(debug, true, (logFile != ""), logFile)
	if serverIsAlive() {
		log.Info("Lantern is up")
	} else {
		log.Info("Lantern is down")
	}
}

func serverStop(cmd *cobra.Command, args []string) {
	log.Init(debug, true, (logFile != ""), logFile)
	if serverIsAlive() {
		log.Info("Lantern will be down")
		stopServer()
	} else {
		log.Info("Lantern is down")
	}
}

func lantern(cmd *cobra.Command, args []string) {
	log.Init(debug, !off, (logFile != ""), logFile)
	runLuz(cmd, args)
}

func runLuz(cmd *cobra.Command, args []string) {
	log.Info("Starting lantern with run command")
	log.DebugWithFields("run called", log.Fields{"args": args})

	// response json
	mapB, _ := json.Marshal(map[string]string{"status": "OK"})
	controller = Controller{cfgFile, string(mapB)}

	// Launch lantern in selected mode
	controller.launchLantern()
	log.Info("End lantern")
}

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

// // //
//
// Mode Server
//
// // //

func (controller Controller) exitHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func (controller Controller) handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, controller.result)
}

func (controller Controller) runServer() {
	http.HandleFunc("/", controller.handler)
	http.HandleFunc("/exit", controller.exitHandler)
	http.ListenAndServe(":8080", nil)
}

func serverIsAlive() bool {
	err := exec.Command("curl", "--silent", "localhost:8080").Run()

	if err != nil {
		return false
	} else {
		return true
	}
}

func stopServer() {
	exec.Command("curl", "--silent", "localhost:8080/exit").Run()
}
