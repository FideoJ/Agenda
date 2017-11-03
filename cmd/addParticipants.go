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

// addParticipantsCmd represents the addParticipants command
var addParticipantsCmd = &cobra.Command{
	Use:   "addParticipants",
	Short: "Add participants to a existed meeting",
	Long: `Add participants to a existed meeting
	args: title (string), participant (string)`,
	Run: func(cmd *cobra.Command, args []string) {
		title := utils.GetNonEmptyString(cmd, "title")
		participants := utils.GetNonEmptyStringSlice(cmd, "participants")

		service.AddParticipants(title, participants)
		logger.Info("AddParticipants called with title: [%+v], participants: [%+v]", title, participants)
	},
}

func init() {
	RootCmd.AddCommand(addParticipantsCmd)

	addParticipantsCmd.Flags().StringP("title", "t", "", "Meeting's title")
	addParticipantsCmd.Flags().StringSliceP("participants", "p", make([]string, 0), "participants' usernames")
}
