package main

import (
	"password/manager/account"
	"password/manager/cipher"
	"password/manager/controller"
	"password/manager/files"
)

func main() {
	var masterPassword string
	isRunning := true
	isLogin := false

	vault := account.NewVault(files.NewJsonDb(account.VaultFileName))

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
