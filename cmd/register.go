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
	"fmt"
	"log"
	"os"

	"github.com/MarshallW906/Agenda/service"
	"github.com/spf13/cobra"
)

var (
	infoLog  *log.Logger
	errorLog *log.Logger
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("register called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")

		infoLog.Printf("register with username:[%+v], password:[%+v], email:[%+v], phone:[%+v]", username, password, email, phone)
		if err := service.Register(username, password, email, phone); err != nil {
			infoLog.Println("Register SUCCEEDED.")
		} else {
			errorLog.Fatalln(err)
		}
	},
}

func init() {
	infoLog = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)

	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "Username")
	registerCmd.Flags().StringP("password", "p", "", "Password")
	registerCmd.Flags().StringP("email", "e", "", "Email")
	registerCmd.Flags().StringP("phone", "t", "", "Phone")
}
