package controller

import (
	"demo/password/account"

	"github.com/fatih/color"
)

func HandleVaultAction(vault *account.Vault, action int8, isRunning *bool) {
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

func createAccount(vault *account.Vault) {
	login := PromptData("Enter login: ")
	password := PromptData("Enter password (or press 'Enter' for generation): ")
	url := PromptData("Enter URL: ")

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
	searchString := PromptData("Enter URL to search: ")

	vault.FindByUrl(searchString)
}

func removeAccount(vault *account.Vault) {
	url := PromptData("Enter URL to remove: ")

	vault.RemoveByUrl(url)
}
