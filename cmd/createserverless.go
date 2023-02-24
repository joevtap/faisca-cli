package cmd

import (
	"fmt"

	"github.com/joevtap/faisca-cli/utils/prompt"
	"github.com/spf13/cobra"
)

var (
	language string
	name     string
)

var createServerlessCmd = &cobra.Command{
	Use:     "serverless [project-name]",
	Short:   "Create a new project using the Serverless Framework",
	Aliases: []string{"sls"},
	Args:    cobra.MaximumNArgs(1),
	Run:     createServerlessCmdImpl,
}

func createServerlessCmdImpl(cmd *cobra.Command, args []string) {
	var err error

	if len(args) != 0 {
		name = args[0]
	}

	if name == "" {
		name, err = prompt.Name.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v \n", err)
			return
		}
	}

	if language == "" {
		_, language, err = prompt.Language.Run([]string{"go", "typescript"})
		if err != nil {
			fmt.Printf("Prompt failed %v \n", err)
			return
		}
	}

	fmt.Printf("Project name: %v\n", name)
	fmt.Printf("Language: %v\n", language)
}

func init() {
	createServerlessCmd.Flags().StringVarP(&language, "lang", "l", "", "Language to use for the project")
	createServerlessCmd.Flags().StringVarP(&name, "name", "", "", "Name of the project")
	createCmd.AddCommand(createServerlessCmd)
}
