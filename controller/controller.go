package controller

import (
	"fmt"

	"github.com/fatih/color"
)

func HandleRegisterMenu() int8 {
	color.Cyan("1. Set up master password")
	color.Cyan("2. Exit")

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
	color.Cyan("1. Login with master password")
	color.Cyan("2. Restart Vault")
	color.Cyan("3. Exit")

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
	color.Cyan("1. Add account")
	color.Cyan("2. Find account")
	color.Cyan("3. Remove account")
	color.Cyan("4. Exit")

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
