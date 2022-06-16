/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"cli/code"
	"cli/helpers"
	"cli/internal"
	template "cli/internal/template"
	"fmt"

	"github.com/spf13/cobra"
)

// ctlCmd represents the ctl command
var ctlCmd = &cobra.Command{
	Use:   "ctl",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		internalRepo := internal.New()
		if len(args) == 0 {
			fmt.Println("Ejemplo: cli mk rts user category")
			return
		}

		module, name := args[0], args[1]

		internalRepo.DirectoryRoot(fmt.Sprintf("%s/%s/%s", code.DIR_MODULES, module, code.DIR_CONTROLLER)).
			CreateFile(fmt.Sprintf("%s.go", helpers.FileName(name))).
			AppendTemplate(template.NewTemplate().ModuleController(name))

		fmt.Println("mk created makefile")
	},
}

func init() {
	mkCmd.AddCommand(ctlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ctlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ctlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
