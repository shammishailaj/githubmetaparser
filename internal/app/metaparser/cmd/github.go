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
	"github.com/shammishailaj/metaparser/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cleanCmd represents the cleanCmd command
var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Parses github meta data",
	Long: `Can be used to retrieve mete data for github`,
	Run: func(cmd *cobra.Command, args []string) {
		appendDeny, appendDenyErr := cmd.Flags().GetBool("deny")
		if appendDenyErr != nil {
			appendDeny = false
		}

		u := utils.NewUtils(&log.Logger{})

		ipv4, ipv4Err := cmd.Flags().GetBool("ipv4")
		if ipv4Err == nil {
			if ipv4 {
				u.PrintGithubIPs(appendDeny)
			}
		}

		ipv6, ipv6Err := cmd.Flags().GetBool("ipv6")
		if ipv6Err == nil {
			if ipv6 {
				fmt.Println("Github does not provide IPv6 IPs")
			}
		}
	},
}



func init() {
	rootCmd.AddCommand(githubCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// githubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	githubCmd.Flags().BoolP("ipv4", "i", false, "Prints List of IPv4 IPs")
	githubCmd.Flags().BoolP("ipv6", "j", false, "Prints List of IPs")
	githubCmd.Flags().BoolP("deny", "d", false, "Append a deny statement to the output")
}

