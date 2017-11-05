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
	"github.com/FideoJ/Agenda/logger"
	"github.com/FideoJ/Agenda/service"

	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout",
	Long: `Logout
	- 用户登出
	- args: None
	- notes: 若未登录，则静默
	`,
	Run: func(cmd *cobra.Command, args []string) {
		service.Logout()

		logger.Info("Logout called")
	},
}

func init() {
	RootCmd.AddCommand(logoutCmd)
}
