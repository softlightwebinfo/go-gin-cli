package cmd

import (
	"cli/code"
	"cli/internal"
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ProcessCreate(args)
	},
}

func ProcessCreate(args []string) {
	if len(args) == 0 {
		fmt.Println("No se ha podido crear el proyecto, necesitas un argumento, ejemplo: cli start appName")
		return
	}

	var nameProject = args[0]
	internalRepo := internal.New()
	internalRepo.CreateDir(nameProject)

	directoryRoot := internalRepo.DirectoryRoot(nameProject).
		CreateDir(code.DIR_CONFIG).
		CreateDir(code.DIR_MODELS).
		CreateDir(code.DIR_MODULES).
		CreateDir(code.DIR_REPOSITORY)

	directoryRoot.CreateFile(code.FILE_ROOT)
	directoryRoot.CreateFile(code.FILE_ENV)
	directoryRoot.CreateFile(code.FILE_ENV_EXAMPLE)
	directoryRoot.CreateFile(code.FILE_GIT_IGNORE)
	directoryRoot.CreateFile(code.FILE_GO_MOD)
}

func init() {
	rootCmd.AddCommand(startCmd)
}
