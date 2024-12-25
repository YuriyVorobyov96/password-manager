package controller

import (
	"password/manager/account"
	"password/manager/cipher"
	"password/manager/files"

	"github.com/fatih/color"
)

func HandleLoginAction(vault *account.VaultWithDb, db *files.MpDb, action int8, masterPassword *string, isLogin *bool, isRunning *bool) {
	switch {
	case action == 1:
		login(db, isLogin, masterPassword)
	case action == 2:
		restartVault(vault, db)
	case action == 3:
		*isRunning = false
	}
}

func login(db *files.MpDb, isLogin *bool, masterPassword *string) {
	enteredMasterPassword := PromptData("Enter master password: ")

	isMatch := cipher.CheckMasterPassword(db, enteredMasterPassword)

	if isMatch {
		color.Green("Correct password. Login...")

		*isLogin = true
		*masterPassword = enteredMasterPassword

		return
	}

	color.Red("Incorrect password")
}

func restartVault(vault *account.VaultWithDb, db *files.MpDb) {
	isRestart := PromptData("Are you sure? This will delete all your data (Y/N): ")

	if isRestart == "y" || isRestart == "Y" {
		vault.Restart()
		cipher.ResetMasterPassword(db)
	}

	if isRestart == "n" || isRestart == "N" {
		HandleLoginMenu()
	}
}
