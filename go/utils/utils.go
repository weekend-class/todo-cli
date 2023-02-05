package utils

import (
	"fmt"

	"todo-cli/constants"
)

func CheckValidCommand(cmd string) (bool, error) {
	for _, c := range constants.AvailableCommands {
		if c == cmd {
			return true, nil
		}
	}

	return false, fmt.Errorf("command not found")
}

func CheckIndexExistsInSlice(arrLength int, index int) bool {
	return index >= 0 && arrLength > index
}
