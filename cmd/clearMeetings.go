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
	"github.com/spf13/cobra"
)

// clearMeetingsCmd represents the clearMeetings command
var clearMeetingsCmd = &cobra.Command{
	Use:   "clearMeetings",
	Short: "Clear all meetings whose sponsor is current user",
	Long: `Clear all meetings whose sponsor is current user
	- 清空会议
	- args: None
	- notes: 要求已登录,清除当前用户为发起者的会议
	`,
	Run: func(cmd *cobra.Command, args []string) {
		service.ClearMeetings()

		logger.Info("ClearMeetings called")
	},
}

func init() {
	RootCmd.AddCommand(clearMeetingsCmd)
}
