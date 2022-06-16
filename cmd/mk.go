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
	"cli/internal"
	template "cli/internal/template"
	"fmt"

	"github.com/spf13/cobra"
)

// mkCmd represents the mk command
var mkCmd = &cobra.Command{
	Use:   "mk",
	Short: "Create makefile files",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		internalRepo := internal.New()

		tog, _ := cmd.Flags().GetBool("makefile")

		if tog {
			internalRepo.DirectoryRoot(".").CreateFile(code.FILE_MAKE).AppendTemplate(template.NewTemplate().Makefile())
			fmt.Println("mk created makefile")
		}
	},
}

func init() {
	mkCmd.Flags().BoolP("makefile", "m", false, "Create makefile")
	rootCmd.AddCommand(mkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
