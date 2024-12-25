package cipher

import (
	"password/manager/files"
	"password/manager/output"

	"github.com/fatih/color"
)

const MasterPasswordFileName = "mp.dat"

func IsMasterPasswordExist(db *files.MpVault) bool {
	_, err := db.Read()

	return err == nil
}

func CreateMasterPassword(db *files.MpVault, password string) {
	if len(password) < 10 {
		output.PrintError("The length of the password must be greater than or equal to 10 characters")

		return
	}

	hashedPassword, err := Hash(password)

	if err != nil {
		output.PrintError("Can't store master password")
	}

	db.Write([]byte(hashedPassword))
	color.Green("Successfully add master password")
}

func CheckMasterPassword(db *files.MpVault, password string) bool {
	hash, err := db.Read()

	if err != nil {
		output.PrintError("Master password data is broken. Please restart Vault")
	}

	return CheckHash(password, string(hash))
}

func ResetMasterPassword(db *files.MpVault) {
	db.Remove()
}
