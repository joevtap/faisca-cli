package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Used to create a new project",
	Long: `The "create" command is used to create a new project.

You can use it to create a project using the predefined languages and frameworks.
`,
	Example: `  faisca create serverless`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
