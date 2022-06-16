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

// midCmd represents the mid command
var midCmd = &cobra.Command{
	Use:   "mid",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		internalRepo := internal.New()

		if len(args) == 0 {
			fmt.Println("Ejemplo: cli mk mid cors")
			return
		}

		internalRepo.
			DirectoryRoot(code.DIR_MIDDLEWARE).
			CreateFile(fmt.Sprintf("%s.go", helpers.FileName(args[0]))).
			AppendTemplate(template.NewTemplate().MiddlewareEmpty(args[0]))

		fmt.Println("middl created")
	},
}

func init() {
	mkCmd.AddCommand(midCmd)
}
