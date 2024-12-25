package controller

import (
	"fmt"
	"password/manager/output"
)

func HandleRegisterMenu() int8 {
	output.PrintAction("1. Set up master password")
	output.PrintAction("2. Exit")

	var res int8

	fmt.Print("Enter operation number: ")
	_, err := fmt.Scan(&res)

	if err != nil || res < 1 || res > 2 {
		fmt.Println()
		return HandleRegisterMenu()
	}

	return res
}

func HandleLoginMenu() int8 {
	output.PrintAction("1. Login with master password")
	output.PrintAction("2. Restart Vault")
	output.PrintAction("3. Exit")

	var res int8

	fmt.Print("Enter operation number: ")
	_, err := fmt.Scan(&res)

	if err != nil || res < 1 || res > 3 {
		fmt.Println()
		return HandleLoginMenu()
	}

	return res
}

func HandleVaultMenu() int8 {
	output.PrintAction("1. Add account")
	output.PrintAction("2. Find account")
	output.PrintAction("3. Remove account")
	output.PrintAction("4. Exit")

	var res int8

	fmt.Print("Enter operation number: ")
	_, err := fmt.Scan(&res)

	if err != nil || res < 1 || res > 4 {
		fmt.Println()
		return HandleVaultMenu()
	}

	return res
}

func PromptData(prompt string) string {
	var res string

	fmt.Print(prompt)
	fmt.Scanln(&res)

	return res
}
