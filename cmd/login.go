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
	"github.com/MarshallW906/Agenda/logger"
	"github.com/MarshallW906/Agenda/service"
	"github.com/MarshallW906/Agenda/utils"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login",
	Long:  `Login`,
	Run: func(cmd *cobra.Command, args []string) {
		username := utils.GetNonEmptyString(cmd, "username")
		password := utils.GetNonEmptyString(cmd, "password")

		service.Login(username, password)

		logger.Info("Login called with username:[%+v], password:[%+v]", username, password)
	},
}

func init() {
	loginCmd.Flags().StringP("username", "u", "", "Username")
	loginCmd.Flags().StringP("password", "p", "", "Password")

	RootCmd.AddCommand(loginCmd)
}
