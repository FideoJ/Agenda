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
	"github.com/FideoJ/Agenda/utils"
	"github.com/spf13/cobra"
)

// removeParticipantsCmd represents the removeParticipants command
var removeParticipantsCmd = &cobra.Command{
	Use:   "removeParticipants",
	Short: "Remove participants of a existed meeting",
	Long: `Remove participants of a existed meeting
	- usage: 删除会议参与者
	- args: title string, participant string
	- notes: 要求已登录,仅能操作当前用户为发起者的会议，仅剩发起者的会议应删除
	`,
	Run: func(cmd *cobra.Command, args []string) {
		title := utils.GetNonEmptyString(cmd, "title")
		participants := utils.GetNonEmptyStringSlice(cmd, "participants")

		service.RemoveParticipants(title, participants)
		logger.Info("RemoveParticipants called with title: [%+v], participants: [%+v]", title, participants)
	},
}

func init() {
	RootCmd.AddCommand(removeParticipantsCmd)

	removeParticipantsCmd.Flags().StringP("title", "t", "", "meeting's title")
	removeParticipantsCmd.Flags().StringSliceP("participants", "p", make([]string, 0), "participants to remove")
}
