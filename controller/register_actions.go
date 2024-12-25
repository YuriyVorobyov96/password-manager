package controller

import (
	"password/manager/cipher"
	"password/manager/files"
)

func HandleRegisterAction(db *files.MpVault, action int8, isRunning *bool) {
	switch {
	case action == 1:
		register(db)
	case action == 2:
		*isRunning = false
	}
}

func register(db *files.MpVault) {
	masterPassword := PromptData("Enter master password: ")

	cipher.CreateMasterPassword(db, masterPassword)
}
