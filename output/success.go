package output

import "github.com/fatih/color"

func PrintSuccess(value string, formatting ...any) {
	color.Green(value, formatting...)
}
