/*
Copyright © 2020 Shammi Shailaj <shammishailaj@gmail.com>

Licensed under the MIT License (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/shammishailaj/metaparser/blob/main/LICENSE

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/shammishailaj/metaparser/pkg/utils"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "metaparser",
		Short: "The Metaparser CLI",
		Long: `The Metaparser CLI parses meta information to extract IPs/ CIDRs from Github, Cloudflare and AWS (and more) for whitelisting. The output is produced in an nginx compatible format so you can use the binary executable in your cron to automatically refresh the list of IPs`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
	}
	logger *log.Logger
	u *utils.Utils
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	logger = log.New()
	u = utils.NewUtils(logger)
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.obcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//rootCmd.Flags().BoolP("ipv4", "i", false, "Prints List of IPv4 IPs")
	//rootCmd.Flags().BoolP("ipv6", "j", false, "Prints List of IPs")
	//rootCmd.Flags().BoolP("deny", "d", false, "Append a deny statement to the output")
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

		// Search config in home directory with name ".obcli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".obcli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
