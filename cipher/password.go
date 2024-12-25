package cipher

import (
	"password/manager/files"

	"github.com/fatih/color"
)

const MasterPasswordFileName = "mp.dat"

func IsMasterPasswordExist(db *files.MpDb) bool {
	_, err := db.ReadFile()

	return err == nil
}

func CreateMasterPassword(db *files.MpDb, password string) {
	if len(password) < 10 {
		color.Red("The length of the password must be greater than or equal to 10 characters")

		return
	}

	hashedPassword, err := Hash(password)

	if err != nil {
		color.Red("Can't store master password")
	}

	db.WriteFile([]byte(hashedPassword))
	color.Green("Successfully add master password")
}

func CheckMasterPassword(db *files.MpDb, password string) bool {
	hash, err := db.ReadFile()

	if err != nil {
		color.Red("Master password data is broken. Please restart Vault")
	}

	return CheckHash(password, string(hash))
}

func ResetMasterPassword(db *files.MpDb) {
	db.RemoveFile()
}
