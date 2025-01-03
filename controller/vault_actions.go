package controller

import (
	"password/manager/account"
	"password/manager/output"
)

func HandleVaultAction(vault *account.VaultWithDb, action int8, masterPassword string, isRunning *bool) {
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

func createAccount(vault *account.VaultWithDb, masterPassword string) {
	login := PromptData("Enter login: ")
	password := PromptData("Enter password (or press 'Enter' for generation): ")
	url := PromptData("Enter URL: ")

	acc, err := account.NewAccount(login, password, url, masterPassword)

	if err != nil {
		if err.Error() == "INVALID_LOGIN" {
			output.PrintError("Invalid login format")

			return
		}
		if err.Error() == "INVALID_URL" {
			output.PrintError("Invalid URL format")

			return
		}
		if err.Error() == "INVALID_PASSWORD" {
			output.PrintError("Invalid password format")

			return
		}
	}

	vault.AddAccount(*acc)
}

func findByUrl(vault *account.VaultWithDb, masterPassword string) {
	searchString := PromptData("Enter URL to search: ")

	vault.FindByUrl(searchString, masterPassword)
}

func removeAccount(vault *account.VaultWithDb) {
	url := PromptData("Enter URL to remove: ")

	vault.RemoveByUrl(url)
}
