package files

import (
	"os"
	"password/manager/output"
)

func readFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func writeFile(content []byte, filename string) {
	file, err := os.Create(filename)

	if err != nil {
		output.PrintError(err)

		return
	}

	_, err = file.Write(content)

	defer file.Close()

	if err != nil {
		output.PrintError(err)

		return
	}
}

func removeFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		output.PrintError(err)
	}
}
