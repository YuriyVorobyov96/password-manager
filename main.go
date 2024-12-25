package main

import (
	"demo/password/account"
	"demo/password/cipher"
	"demo/password/controller"
)

func main() {
	isRunning := true
	isLogin := false
	vault := account.NewVault()

	for isRunning {
		if cipher.IsMasterPasswordExist() {
			if isLogin {
				action := controller.HandleVaultMenu()
				controller.HandleVaultAction(vault, action, &isRunning)
				continue
			}

			action := controller.HandleLoginMenu()
			controller.HandleLoginAction(vault, action, &isLogin, &isRunning)
			continue
		}

		action := controller.HandleRegisterMenu()
		controller.HandleRegisterAction(action, &isRunning)
	}
}
