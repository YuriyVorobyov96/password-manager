package files

import (
	"os"

	"github.com/fatih/color"
)

func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)

	if err != nil {
		color.Red(err.Error())

		return
	}

	_, err = file.Write(content)

	defer file.Close()

	if err != nil {
		color.Red(err.Error())

		return
	}
}

func RemoveFile(name string) {
	err := os.Remove(name)

	if err != nil {
		color.Red(err.Error())
	}
}
