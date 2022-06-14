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
	"fmt"
	"regexp"
	"strings"

	template "cli/internal/template"

	"github.com/spf13/cobra"
)

// libCmd represents the lib command
var libCmd = &cobra.Command{
	Use:   "lib",
	Short: "Library functions",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			return
		}

		nameLibrary := args[0]
		var filenames []string

		dirLibrary := internal.NewDirectoryRoot(code.DIR_LIBRARIES)
		r, _ := regexp.Compile("(^[a-z]+[a-z])|([A-Z]+[a-z]*)")

		for _, item := range r.FindAllStringSubmatch(nameLibrary, -1) {
			name := item[0]
			filenames = append(filenames, strings.ToLower(name))
		}

		dirLibrary.
			CreateFile(fmt.Sprintf("%s.go", strings.Join(filenames, "_"))).
			AppendTemplate(
				template.NewTemplate().LibraryNew(nameLibrary),
			)
	},
}

func init() {
	mkCmd.AddCommand(libCmd)
}
