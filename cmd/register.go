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
	"../logger"
	"../service"
	"../utils"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a user",
	Long: `Register a user
	- usage: 用户注册
	- args: username string, password string, email string, phone string
	`,
	Run: func(cmd *cobra.Command, args []string) {
		username := utils.GetNonEmptyString(cmd, "username")
		password := utils.GetNonEmptyString(cmd, "password")
		email := utils.GetNonEmptyString(cmd, "email")
		phone := utils.GetNonEmptyString(cmd, "phone")

		service.Register(username, password, email, phone)

		logger.Info("Register called with username:[%+v], password:[%+v], email:[%+v], phone:[%+v]", username, password, email, phone)
	},
}

func init() {
	registerCmd.Flags().StringP("username", "u", "", "Username")
	registerCmd.Flags().StringP("password", "p", "", "Password")
	registerCmd.Flags().StringP("email", "e", "", "Email")
	registerCmd.Flags().StringP("phone", "t", "", "Phone")

	RootCmd.AddCommand(registerCmd)
}
