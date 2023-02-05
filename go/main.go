package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"todo-cli/commands"
	"todo-cli/interfaces"
	"todo-cli/utils"
)

func main() {
	var todos []interfaces.Todo

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		cmd := strings.Split(scanner.Text(), " ")

		_, err := utils.CheckValidCommand(cmd[0])

		if err != nil {
			fmt.Println(err.Error())
			break
		}

		switch cmd[0] {
		case "add":
			todoName := strings.Join(cmd[1:], " ")

			err := commands.Add(&todos, todoName)

			if err != nil {
				fmt.Println(err.Error())
			}
		case "check":
			err := commands.Check(&todos, cmd[1])

			if err != nil {
				fmt.Println(err.Error())
			}
		case "uncheck":
			err := commands.Uncheck(&todos, cmd[1])

			if err != nil {
				fmt.Println(err.Error())
			}
		case "ls":
			err := commands.List(&todos)

			if err != nil {
				fmt.Println(err.Error())
			}
		case "rm":
			err := commands.Remove(&todos, cmd[1])

			if err != nil {
				fmt.Println(err.Error())
			}
		case "toggle":
			err := commands.Toggle(&todos, cmd[1])

			if err != nil {
				fmt.Println(err.Error())
			}
		case "search":
			todoName := strings.Join(cmd[1:], " ")

			err := commands.Search(&todos, todoName)

			if err != nil {
				fmt.Println(err.Error())
			}
		default:
			fmt.Println("command not found")
		}
	}
}
