package usecase

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

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

func ChooseMonitorLog() string {
	prompt := promptui.Select{
		Label: "Please choose monitor log level",
		Items: []string{"off", "debug", "info", "trace", "all"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	fmt.Printf("You choose %q\n", result)
	return result
}
