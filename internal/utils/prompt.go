package utils

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func ReadCmdInput(msg string) string {
	templates := &promptui.PromptTemplates{
		Prompt: "{{ . }} ",
		// Valid:   "{{ . | green }} ",
		// Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}
	prompt := promptui.Prompt{
		Label:     msg,
		Templates: templates,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}
