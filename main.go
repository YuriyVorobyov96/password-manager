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
	mpDb := files.NewMpDb(cipher.MasterPasswordFileName) 
	vault := account.NewVault(files.NewJsonDb(account.VaultFileName))

	for isRunning {
		if cipher.IsMasterPasswordExist(mpDb) {
			if isLogin {
				action := controller.HandleVaultMenu()
				controller.HandleVaultAction(vault, action, masterPassword, &isRunning)
				continue
			}

			action := controller.HandleLoginMenu()
			controller.HandleLoginAction(vault, mpDb, action, &masterPassword, &isLogin, &isRunning)
			continue
		}

		action := controller.HandleRegisterMenu()
		controller.HandleRegisterAction(mpDb, action, &isRunning)
	}
}
