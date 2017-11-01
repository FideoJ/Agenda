// // Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// //
// // Licensed under the Apache License, Version 2.0 (the "License");
// // you may not use this file except in compliance with the License.
// // You may obtain a copy of the License at
// //
// //     http://www.apache.org/licenses/LICENSE-2.0
// //
// // Unless required by applicable law or agreed to in writing, software
// // distributed under the License is distributed on an "AS IS" BASIS,
// // WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// // See the License for the specific language governing permissions and
// // limitations under the License.

package cmd

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/FideoJ/Agenda/service"
// 	"github.com/spf13/cobra"
// )

// var (
// 	infoLog  *log.Logger
// 	errorLog *log.Logger
// )

// // createmtCmd represents the createmt command
// var createmtCmd = &cobra.Command{
// 	Use:   "createmt",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("createmt called")

// 		title, _ := cmd.Flags().GetString("title")
// 		startTimeStr, _ := cmd.Flags().GetString("starttime")
// 		endTimeStr, _ := cmd.Flags().GetString("endtime")
// 		participants, _ := cmd.Flags().GetStringSlice("participants")

// 		infoLog.Printf("Title: [%+v], startTimeStr: [%+v], endTimeStr: [%+v], participants: [%+v]\n", title, startTimeStr, endTimeStr, participants)
// 		if err := service.CreateMt(title, startTimeStr, endTimeStr, participants); err == nil {
// 			infoLog.Println("Create Meeting SUCCEEDED")
// 		} else {
// 			errorLog.Fatalln(err)
// 		}
// 	},
// }

// func init() {
// 	infoLog = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
// 	errorLog = log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)

// 	RootCmd.AddCommand(createmtCmd)

// 	createmtCmd.Flags().StringP("title", "t", "", "meeting's title")
// 	createmtCmd.Flags().StringP("starttime", "s", "", "meeting's start time")
// 	createmtCmd.Flags().StringP("endtime", "e", "", "meeting's end time")
// 	createmtCmd.Flags().StringSliceP("participants", "p", make([]string, 0), "meeting's participants. specify repeatedly to specify each participant.")
// }
