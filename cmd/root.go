/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "airdrop",
		Short: "HungryPanda Token Airdrop Service",
		Long:  ``,
	}
	mappings = map[string]string{
		"env":           "ENV",
		"bind":          "BIND_ADDR",
		"db_url":        "PG_URL",
		"priv_key":      "PRIVATE_KEY",
		"rpc":           "RPC_URL",
		"base_url":      "BASE_URL",
		"fb_app_id":     "FACEBOOK_APP_ID",
		"fb_app_secret": "FACEBOOK_APP_KEY",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.airdrop.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	for k, v := range mappings {
		cobra.CheckErr(viper.BindEnv(k, v))
	}
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".airdrop" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".airdrop")
	}

	// load default
	_ = godotenv.Load()
	_ = godotenv.Load(fmt.Sprintf("%s.env", viper.GetString("env")))

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
