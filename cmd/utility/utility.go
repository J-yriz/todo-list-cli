package utility

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func CloseFile(file *os.File) {

	err := file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func CmdInput(quiz string) (string, error) {
	prompt := promptui.Prompt{
		Label: quiz,
	}

	return prompt.Run()
}

func CmdSelect(label string, option []string) (string, error) {

	prompt := promptui.Select{
		Label: label,
		Items: option,
	}

	_, result, err := prompt.Run()

	return result, err

}

func CmdEdit(label string, defaultValue string) (string, error) {

	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}

	return prompt.Run()

}
