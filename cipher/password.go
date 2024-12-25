package cipher

import (
	"password/manager/files"

	"github.com/fatih/color"
)

const masterPasswordFileName = "mp.dat"

func IsMasterPasswordExist() bool {
	db := files.NewMpDb(masterPasswordFileName)
	_, err := db.ReadFile()

	return err == nil
}

func CreateMasterPassword(password string) {
	if len(password) < 10 {
		color.Red("The length of the password must be greater than or equal to 10 characters")

		return
	}

	hashedPassword, err := Hash(password)

	if err != nil {
		color.Red("Can't store master password")
	}

	db := files.NewMpDb(masterPasswordFileName)
	db.WriteFile([]byte(hashedPassword))
	color.Green("Successfully add master password")
}

func CheckMasterPassword(password string) bool {
	db := files.NewMpDb(masterPasswordFileName)
	hash, err := db.ReadFile()

	if err != nil {
		color.Red("Master password data is broken. Please restart Vault")
	}

	return CheckHash(password, string(hash))
}

func ResetMasterPassword() {
	db := files.NewMpDb(masterPasswordFileName)
	db.RemoveFile()
}
