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
	mpVault := files.NewMpVault(cipher.MasterPasswordFileName)
	vault := account.NewVault(files.NewJsonDb(account.VaultFileName))

	for isRunning {
		if cipher.IsMasterPasswordExist(mpVault) {
			if isLogin {
				action := controller.HandleVaultMenu()
				controller.HandleVaultAction(vault, action, masterPassword, &isRunning)
				continue
			}

			action := controller.HandleLoginMenu()
			controller.HandleLoginAction(vault, mpVault, action, &masterPassword, &isLogin, &isRunning)
			continue
		}

		action := controller.HandleRegisterMenu()
		controller.HandleRegisterAction(mpVault, action, &isRunning)
	}
}
