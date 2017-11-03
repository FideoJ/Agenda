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

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "Create a meeting",
	Long: `Create a meeting
	- 创建会议
	- args: title string, startTime string, endTime string, participants []string
	- notes: 要求已登录,时间格式:"YYYY:MM:DD HH:mm"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		title := utils.GetNonEmptyString(cmd, "title")
		startTimeStr := utils.GetNonEmptyString(cmd, "startTime")
		endTimeStr := utils.GetNonEmptyString(cmd, "endTime")
		participants := utils.GetNonEmptyStringSlice(cmd, "participants")

		service.CreateMeeting(title, startTimeStr, endTimeStr, participants)

		logger.Info("CreateMeeting called with title: [%+v], startTime: [%+v], endTime: [%+v], participants: [%+v]\n", title, startTimeStr, endTimeStr, participants)
	},
}

func init() {
	RootCmd.AddCommand(createMeetingCmd)

	createMeetingCmd.Flags().StringP("title", "t", "", "meeting's title")
	createMeetingCmd.Flags().StringP("startTime", "s", "", "meeting's start time")
	createMeetingCmd.Flags().StringP("endTime", "e", "", "meeting's end time")
	createMeetingCmd.Flags().StringSliceP("participants", "p", make([]string, 0), "meeting's participants. specify repeatedly to specify each participant.")
}
