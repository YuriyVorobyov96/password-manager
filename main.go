package main

import (
	"demo/password/account"
	"demo/password/cipher"
	"demo/password/controller"
)

func main() {
	var masterPassword string
	isRunning := true
	isLogin := false
	vault := account.NewVault()

	for isRunning {
		if cipher.IsMasterPasswordExist() {
			if isLogin {
				action := controller.HandleVaultMenu()
				controller.HandleVaultAction(vault, action, masterPassword, &isRunning)
				continue
			}

			action := controller.HandleLoginMenu()
			controller.HandleLoginAction(vault, action, &masterPassword, &isLogin, &isRunning)
			continue
		}

		action := controller.HandleRegisterMenu()
		controller.HandleRegisterAction(action, &isRunning)
	}
}
