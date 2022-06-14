package cmd

import (
	"cli/code"
	"cli/internal"
	template "cli/internal/template"
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
	generate(nameProject)
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func generate(nameProject string) {
	internalRepo := internal.New()
	internalRepo.CreateDir(nameProject)

	directoryRoot := internalRepo.DirectoryRoot(nameProject).
		CreateDir(code.DIR_CONFIG).
		CreateDir(code.DIR_MIDDLEWARE).
		CreateDir(code.DIR_STATIC).
		CreateDir(code.DIR_ROUTER).
		CreateDir(code.DIR_MODELS).
		CreateDir(code.DIR_MODULES).
		CreateDir(code.DIR_LIBRARIES).
		CreateDir(code.DIR_REPOSITORY)

	directoryRoot.CreateFile(code.FILE_ROOT).AppendTemplate(template.NewTemplate().Main(nameProject))
	directoryRoot.CreateFile(code.FILE_ENV).AppendTemplate(template.NewTemplate().Env())
	directoryRoot.CreateFile(code.FILE_ENV_EXAMPLE).AppendTemplate(template.NewTemplate().EnvExample())
	directoryRoot.CreateFile(code.FILE_GIT_IGNORE).AppendTemplate(template.NewTemplate().Gitignore())
	directoryRoot.CreateFile(code.FILE_GO_MOD).AppendTemplate(template.NewTemplate().GoMod(nameProject))

	entryDirectoryConfig := directoryRoot.EntryDirectory(code.DIR_CONFIG)
	entryDirectoryConfig.CreateFile(code.FILE_CONFIG).AppendTemplate(template.NewTemplate().Config())
	entryDirectoryConfig.CreateFile(code.FILE_CONFIG_DB).AppendTemplate(template.NewTemplate().ConfigDB())
	entryDirectoryConfig.CreateFile(code.FILE_CONFIG_EMAIL).AppendTemplate(template.NewTemplate().ConfigEmail())

	entryDirectoryStatic := directoryRoot.EntryDirectory(code.DIR_STATIC)
	entryDirectoryStatic.CreateFile(code.FILE_STATIC).AppendTemplate(template.NewTemplate().Static())

	entryDirectoryMiddleware := directoryRoot.EntryDirectory(code.DIR_MIDDLEWARE)
	entryDirectoryMiddleware.CreateFile(code.FILE_MIDDLEWARE).AppendTemplate(template.NewTemplate().Middleware())

	entryDirectoryRouter := directoryRoot.EntryDirectory(code.DIR_ROUTER)
	entryDirectoryRouter.CreateFile(code.FILE_ROUTER).AppendTemplate(template.NewTemplate().Router())

	entryDirectoryLibraries := directoryRoot.EntryDirectory(code.DIR_LIBRARIES)
	entryDirectoryLibraries.CreateFile(code.FILE_LIBRARY).AppendTemplate(template.NewTemplate().Library())
}
