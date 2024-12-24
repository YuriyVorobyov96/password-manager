package files

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println(err)

		return
	}

	_, err = file.Write(content)

	defer file.Close()

	if err != nil {
		fmt.Println(err)

		return
	}

	color.Green("Success write to file")
}