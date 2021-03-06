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

	"github.com/spf13/cobra"
)

// cfgCmd represents the cfg command
var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		internalRepo := internal.New()
		if len(args) == 0 {
			return
		}

		name := args[0]
		fileName := helpers.FileName(name)

		dirModules := internalRepo.
			DirectoryRoot(code.DIR_MODULES).
			CreateDir(fileName).
			EntryDirectory(fileName)

		dirModules.CreateDir(code.DIR_CONTROLLER)
		dirModules.CreateDir(code.DIR_MODEL)
		dirModules.CreateDir(code.DIR_REPOSITORY)
		dirModules.CreateDir(code.DIR_ROUTER)

		dirModules.CreateFile(code.FILE_REPOSITORY).
			AppendTemplate(template.NewTemplate().ModuleBaseRepository(name))
		dirModules.CreateFile(code.FILE_ROUTER)
	},
}

func init() {
	mkCmd.AddCommand(cfgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cfgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cfgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
