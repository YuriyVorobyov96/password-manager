package cipher

import (
	"demo/password/files"

	"github.com/fatih/color"
)

const masterPasswordFileName = "mp.dat"

func IsMasterPasswordExist() bool {
	_, err := files.ReadFile(masterPasswordFileName)

	return err != nil
}

func CreateMasterPassword(password string) {
	if len(password) < 10 {
		color.Red("The length of the password must be greater than or equal to 10 characters")
	}

	hashedPassword, err := Hash(password)

	if err != nil {
		color.Red("Can't store master password")
	}

	files.WriteFile([]byte(hashedPassword), masterPasswordFileName)
	color.Green("Successfully add master password")
}
