package controller

import (
	"password/manager/account"

	"github.com/fatih/color"
)

func HandleVaultAction(vault *account.Vault, action int8, masterPassword string, isRunning *bool) {
	switch {
	case action == 1:
		createAccount(vault, masterPassword)
	case action == 2:
		findByUrl(vault, masterPassword)
	case action == 3:
		removeAccount(vault)
	case action == 4:
		*isRunning = false
	}
}

func createAccount(vault *account.Vault, masterPassword string) {
	login := PromptData("Enter login: ")
	password := PromptData("Enter password (or press 'Enter' for generation): ")
	url := PromptData("Enter URL: ")

	acc, err := account.NewAccount(login, password, url, masterPassword)

	if err != nil {
		if err.Error() == "INVALID_LOGIN" {
			color.Red("Invalid login format")

			return
		}
		if err.Error() == "INVALID_URL" {
			color.Red("Invalid URL format")

			return
		}
		if err.Error() == "INVALID_PASSWORD" {
			color.Red("Invalid password format")

			return
		}
	}

	vault.AddAccount(*acc)
}

func findByUrl(vault *account.Vault, masterPassword string) {
	searchString := PromptData("Enter URL to search: ")

	vault.FindByUrl(searchString, masterPassword)
}

func removeAccount(vault *account.Vault) {
	url := PromptData("Enter URL to remove: ")

	vault.RemoveByUrl(url)
}
