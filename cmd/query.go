// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query [-uUsername] [-eEmail]",
	Short: "Query a User by Username and/or E-mail.",
	Long: `Query a User by Username and/or E-mail.

If there is a username contains the provided string,
or there is an e-mail contains the provided string,
Agenda query will print that user.
If neither -u nor -e is specified, query will simply print all the users.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("query called")
		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")

		fmt.Printf("query Users by Username:[%+v], by Email:[%+v]\n", username, email)
	},
}

func init() {
	queryCmd.Flags().StringP("username", "u", "", "query string on username")
	queryCmd.Flags().StringP("email", "e", "", "query string on e-mail")

	RootCmd.AddCommand(queryCmd)
}
