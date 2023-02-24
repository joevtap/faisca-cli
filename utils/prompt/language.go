package prompt

import "github.com/manifoldco/promptui"

type language struct{}

var Language language

func (language) Run(options []string) (int, string, error) {
	prompt := promptui.Select{
		Label:    "Select a language",
		Items:    options,
		HideHelp: true,
		Templates: &promptui.SelectTemplates{
			Active:   "🤔 {{ . | cyan }}",
			Inactive: "   {{ . | cyan }}",
			Selected: "😄 {{ . | green }}",
		},
	}

	return prompt.Run()
}
