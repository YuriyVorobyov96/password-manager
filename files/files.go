package files

import (
	"os"

	"github.com/fatih/color"
)

func readFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func writeFile(content []byte, filename string) {
	file, err := os.Create(filename)

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

func removeFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		color.Red(err.Error())
	}
}
