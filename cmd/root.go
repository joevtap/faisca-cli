package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "faisca",
	Short: "A polyglot project starter and manager",
	Long: `faisca is a polyglot project starter and manager. It is a CLI tool that
allows you to create new projects from templates, and manage existing projects.
`, // + "It is designed to be extensible, so that you can add your own templates and commands."
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
