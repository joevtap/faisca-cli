package prompt

import (
	"fmt"
	"regexp"

	"github.com/manifoldco/promptui"
)

type name struct{}

var Name name

func (name) Run() (string, error) {
	prompt := promptui.Prompt{
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

	return prompt.Run()
}
