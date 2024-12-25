package controller

import (
	"demo/password/account"
	"demo/password/cipher"

	"github.com/fatih/color"
)

func HandleLoginAction(vault *account.Vault, action int8, isLogin *bool, isRunning *bool) {
	switch {
	case action == 1:
		login(isLogin)
	case action == 2:
		restartVault(vault)
	case action == 3:
		*isRunning = false
	}
}

func login(isLogin *bool) {
	masterPassword := PromptData("Enter master password: ")

	isMatch := cipher.CheckMasterPassword(masterPassword)

	if isMatch {
		color.Green("Correct password. Login...")

		*isLogin = true
		return
	}

	color.Red("Incorrect password")
}

func restartVault(vault *account.Vault) {
	isRestart := PromptData("Are you sure? This will delete all your data (Y/N): ")

	if isRestart == "y" || isRestart == "Y" {
		vault.Restart()
		cipher.ResetMasterPassword()
	}

	if isRestart == "n" || isRestart == "N" {
		HandleLoginMenu()
	}
}
