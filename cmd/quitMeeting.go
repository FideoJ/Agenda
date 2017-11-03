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

// quitMeetingCmd represents the quitMeeting command
var quitMeetingCmd = &cobra.Command{
	Use:   "quitMeeting",
	Short: "Quit a meeting",
	Long: `Quit a meeting
	- 退出会议
	- args: title string
	- notes: 要求已登录,仅能操作当前用户为参与者的会议
	`,
	Run: func(cmd *cobra.Command, args []string) {
		title := utils.GetNonEmptyString(cmd, "title")

		service.QuitMeeting(title)

		logger.Info("QuitMeeting called with title: [%+v]", title)
	},
}

func init() {
	RootCmd.AddCommand(quitMeetingCmd)

	quitMeetingCmd.Flags().StringP("title", "t", "", "meeting's title")
}
