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

// removeParticipantCmd represents the removeParticipant command
var removeParticipantCmd = &cobra.Command{
	Use:   "removeParticipant",
	Short: "remove a participant of a existed meeting",
	Long:  `remove a participant of a existed meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title := utils.GetNonEmptyString(cmd, "title")
		participants := utils.GetNonEmptyStringSlice(cmd, "participants")

		service.RemoveParticipant(title, participants)
		logger.Info("removeParticipant called with title: [%+v], participants: [%+v]", title, participants)
	},
}

func init() {
	RootCmd.AddCommand(removeParticipantCmd)

	removeParticipantCmd.Flags().StringP("title", "t", "", "meeting's title")
	removeParticipantCmd.Flags().StringSliceP("participants", "p", make([]string, 0), "participants to remove")
}
