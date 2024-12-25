package controller

import (
	"password/manager/cipher"
	"password/manager/files"
)

func HandleRegisterAction(db *files.MpDb, action int8, isRunning *bool) {
	switch {
	case action == 1:
		register(db)
	case action == 2:
		*isRunning = false
	}
}

func register(db *files.MpDb) {
	masterPassword := PromptData("Enter master password: ")

	cipher.CreateMasterPassword(db, masterPassword)
}
