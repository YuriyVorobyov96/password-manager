package controller

import (
	"demo/password/cipher"
)

func HandleRegisterAction(action int8, isRunning *bool) {
	switch {
	case action == 1:
		register()
	case action == 2:
		*isRunning = false
	}
}

func register() {
	masterPassword := PromptData("Enter master password: ")

	cipher.CreateMasterPassword(masterPassword)
}
