package main

import (
	"demo/password/account"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	isRunning := true
	vault := account.NewVault()

	for isRunning {
		action := handleMenu()
		handleAction(vault, action, &isRunning)
	}
}

func handleMenu() int8 {
	color.Cyan("1. Add account")
	color.Cyan("2. Find account")
	color.Cyan("3. Remove account")
	color.Cyan("4. Exit")

	var res int8

	fmt.Print("Enter operation number: ")
	_, err := fmt.Scan(&res)

	if err != nil || res < 1 || res > 4 {
		fmt.Println()
		return handleMenu()
	}

	return res
}

func handleAction(vault *account.Vault, action int8, isRunning *bool) {
	switch {
	case action == 1:
		createAccount(vault)
	case action == 2:
		findByUrl(vault)
	case action == 3:
		removeAccount(vault)
	case action == 4:
		*isRunning = false
	}
}

func promptData(prompt string) string {
	var res string

	fmt.Print(prompt)
	fmt.Scanln(&res)

	return res
}

func createAccount(vault *account.Vault) {
	login := promptData("Enter login: ")
	password := promptData("Enter password: ")
	url := promptData("Enter URL: ")

	acc, err := account.NewAccount(login, password, url)

	if err != nil {
		if err.Error() == "INVALID_LOGIN" {
			color.Red("Invalid login format")

			return
		}
		if err.Error() == "INVALID_URL" {
			color.Red("Invalid URL format")

			return
		}
	}

	vault.AddAccount(*acc)
}

func findByUrl(vault *account.Vault) {
	searchString := promptData("Enter URL to search: ")

	vault.FindByUrl(searchString)
}

func removeAccount(vault *account.Vault) {
	url := promptData("Enter URL to remove: ")

	vault.RemoveByUrl(url)
}
