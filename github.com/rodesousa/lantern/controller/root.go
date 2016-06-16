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
	"github.com/rodesousa/lantern/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
	logFile string
	debug   bool
	off     bool
	server  bool
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

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		RootCmd.SetOutput(logger.GetOutLogger())
		RootCmd.Help()
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initFromCL)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&logFile, "logfile", "", "log file output (default is current path)")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "show debug message")
	RootCmd.PersistentFlags().BoolVarP(&off, "off", "o", false, "disable out console log")
	RootCmd.PersistentFlags().BoolVarP(&server, "server", "s", false, "mode lantern server")
	//RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.luz-lantern.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initFromCL() {

	//mode server, disabled out console log
	if server {
		off = true
	}

	// initialize debug level
	logger.Init(debug, !off, (logFile != ""), logFile)

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
