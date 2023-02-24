package cmd

import (
	"fmt"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	use            string
	template       string
	availableLangs = []string{"go"}
)

var createServerlessCmd = &cobra.Command{
	Use:     "serverless [project-name]",
	Short:   "Create a new project using the Serverless Framework",
	Args:    cobra.MaximumNArgs(1),
	Aliases: []string{"sls"},
	Run:     createServerlessCmdImpl,
}

func createServerlessCmdImpl(cmd *cobra.Command, args []string) {
	var err error
	var name string

	if len(args) != 0 {
		name = args[0]
	} else {
		name, err = namePrompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v \n", err)
		}
	}

	languagePrompt.Run()
	fmt.Printf("Project name: %v\n", name)
}

func init() {
	createServerlessCmd.Flags().StringVarP(&use, "use", "u", "", "Use a language or framework")
	createServerlessCmd.Flags().StringVarP(&template, "template", "t", "", "Use a specific template (git repo)")
	createCmd.AddCommand(createServerlessCmd)
}

var namePrompt = promptui.Prompt{
	Label:   "Give your project a name",
	Default: "my-project",
	Validate: func(input string) error {
		if len(input) < 3 {
			return fmt.Errorf("project name must be at least 3 characters")
		}

		if !regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString(input) {
			return fmt.Errorf("project name must be alphanumeric")
		}

		return nil
	},
}

var languagePrompt = promptui.Select{
	Label:    "Select a language",
	Items:    availableLangs,
	HideHelp: true,
	Templates: &promptui.SelectTemplates{
		Active:   "🤔 {{ . | cyan }}",
		Inactive: "   {{ . | cyan }}",
		Selected: "😄 {{ . | green }}",
	},
}
