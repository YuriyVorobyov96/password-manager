package main

import (
	"demo/password/account"
	"demo/password/controller"
)

func main() {
	isRunning := true
	vault := account.NewVault()

	for isRunning {
		action := controller.HandleMenu()
		controller.HandleAction(vault, action, &isRunning)
	}
}
