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

// cancelMeetingCmd represents the cancelMeeting command
var cancelMeetingCmd = &cobra.Command{
	Use:   "cancelMeeting",
	Short: "Cancel a meeting",
	Long: `Cancel a meeting
	- 取消会议
	- args: title (string)
	- notes: 要求已登录,仅能操作当前用户为发起者的会议
	`,
	Run: func(cmd *cobra.Command, args []string) {
		title := utils.GetNonEmptyString(cmd, "title")

		service.CancelMeeting(title)

		logger.Info("CancelMeeting called with title: [%+v]", title)
	},
}

func init() {
	RootCmd.AddCommand(cancelMeetingCmd)

	cancelMeetingCmd.Flags().StringP("title", "t", "", "meeting's title")
}
