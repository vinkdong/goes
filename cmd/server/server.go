// Copyright © 2018 VinkDong <dong@wenqi.us>
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

package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vinkdong/goes/pkg/server"
	"github.com/vinkdong/gox/log"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var serverCmd = &cobra.Command{
	Use:   "goes",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		address, err := cmd.Flags().GetString("address")
		next, err := cmd.Flags().GetString("redirect")
		if redirect := os.Getenv("REDIRECT_TO"); redirect != "" {
			next = redirect
		}
		version := os.Getenv("VERSION")
		if err != nil {
			log.Info(err)
		}
		s := server.NewServer()
		s.Next = next
		s.Addr = address
		s.Version = version
		err = s.Run()
		if err != nil {
			log.Error(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := serverCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goes.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	serverCmd.Flags().StringP("address", "a", ":9000", "Help message for toggle")
	serverCmd.Flags().StringP("redirect","r","","where redirect request to")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".goes" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".goes")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	serverCmd.Execute()
}
