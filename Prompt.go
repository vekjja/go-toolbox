package toolbox

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

// promptConfirm displays a Yes/No choice and returns true if user selects "Yes".
func PromptConfirm(label string) bool {
	prompt := promptui.Select{
		Label: label,
		Items: []string{"No", "Yes"},
	}

	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return false
	}
	// "Yes" is index 1 in the slice above
	return i == 1
}

// promptInput shows a single-line prompt and returns the user input.
func PromptInput(label string, defaultVal string) string {
	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultVal,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

// promptSelect builds a simple list-based selector from a slice of strings.
func PromptSelect(label string, items []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	i, _, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return items[i]
}
