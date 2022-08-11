package usecase

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type PromptMessage struct {
	Label  string
	ErrMsg string
}

// func ChooseFabricVersion(pm PromptMessage) string {
// 	validate := func(input string) error {
// 		_, err := strconv.ParseFloat(input, 64)
// 		if err != nil {
// 			return errors.New(pm.ErrMsg)
// 		}
// 		return nil
// 	}
// 	templates := &promptui.PromptTemplates{
// 		Prompt:  "{{ . }} ",
// 		Valid:   "{{ . | green }} ",
// 		Invalid: "{{ . | red }} ",
// 		Success: "{{ . | bold }} ",
// 	}
// 	prompt := promptui.Prompt{
// 		Label:     pm.Label,
// 		Validate:  validate,
// 		Templates: templates,
// 	}

// 	result, err := prompt.Run()

// 	if err != nil {
// 		fmt.Printf("Prompt failed %v\n", err)
// 		os.Exit(1)
// 	}

// 	fmt.Printf("You choose %q\n", result)
// 	return result
// }
func ChooseFabricVersion() string {
	prompt := promptui.Select{
		Label: "Please choose fabric version",
		Items: []string{"1.4.6", "2.2.4"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	fmt.Printf("You choose %q\n", result)
	return result
}
