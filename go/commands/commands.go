package commands

import (
	"fmt"
	"strconv"
	"strings"

	"todo-cli/interfaces"
	"todo-cli/utils"
)

func Add(todos *[]interfaces.Todo, todoName string) error {
	if todoName == "" {
		return fmt.Errorf("todo cannot be empty")
	}

	todo := interfaces.Todo{
		Name:      todoName,
		Completed: false,
	}

	*todos = append(*todos, todo)

	fmt.Printf("New todo (%s) has been added ✔️\n", todo.Name)

	return nil
}

func Check(todos *[]interfaces.Todo, rawIdx string) error {
	if rawIdx == "" {
		return fmt.Errorf("the input has to be the index of todo")
	}

	idx, err := strconv.ParseInt(rawIdx, 10, 64)

	if err != nil {
		return fmt.Errorf("the input has to be an integer")
	}

	parsedIdx := idx - 1

	if !utils.CheckIndexExistsInSlice(len(*todos), int(parsedIdx)) {
		return fmt.Errorf("todo index is not found")
	}

	temp := (*todos)[parsedIdx]

	if temp.Completed {
		return fmt.Errorf("the todo is already completed, no need to remark it again")
	}

	(*todos)[parsedIdx] = interfaces.Todo{
		Name:      (*todos)[parsedIdx].Name,
		Completed: true,
	}

	fmt.Printf("%s already completed ✔️\n", temp.Name)

	return nil
}

func Uncheck(todos *[]interfaces.Todo, rawIdx string) error {
	if rawIdx == "" {
		return fmt.Errorf("the input has to be the index of todo")
	}

	idx, err := strconv.ParseInt(rawIdx, 10, 64)

	if err != nil {
		return fmt.Errorf("the input has to be an integer")
	}

	parsedIdx := idx - 1

	if !utils.CheckIndexExistsInSlice(len(*todos), int(parsedIdx)) {
		return fmt.Errorf("todo index is not found")
	}

	temp := (*todos)[parsedIdx]

	if !temp.Completed {
		return fmt.Errorf("the todo is not completed yet, no need to remark it again")
	}

	(*todos)[parsedIdx] = interfaces.Todo{
		Name:      (*todos)[parsedIdx].Name,
		Completed: false,
	}

	fmt.Printf("%s already marked as uncompleted ✔️\n", temp.Name)

	return nil
}

func List(todos *[]interfaces.Todo) error {
	if len(*todos) <= 0 {
		return fmt.Errorf("no todo available")

	}

	for i, todo := range *todos {
		var checkMark string

		if todo.Completed {
			checkMark = " ✔️"
		}

		fmt.Printf("%v) %s %s\n", i+1, todo.Name, checkMark)
	}

	return nil
}

func Remove(todos *[]interfaces.Todo, rawIdx string) error {
	if rawIdx == "" {
		return fmt.Errorf("the input has to be the index of todo")
	}

	idx, err := strconv.ParseInt(rawIdx, 10, 64)

	if err != nil {
		return fmt.Errorf("the input has to be an integer")
	}

	parsedIdx := idx - 1

	if !utils.CheckIndexExistsInSlice(len(*todos), int(parsedIdx)) {
		return fmt.Errorf("todo index is not found")
	}

	temp := (*todos)[parsedIdx]

	*todos = append((*todos)[:parsedIdx], (*todos)[parsedIdx+1:]...)

	fmt.Printf("%v already removed from the todo list ✔️\n", temp.Name)

	return nil
}

func Toggle(todos *[]interfaces.Todo, rawIdx string) error {
	if rawIdx == "" {
		return fmt.Errorf("the input has to be the index of todo")
	}

	idx, err := strconv.ParseInt(rawIdx, 10, 64)

	if err != nil {
		return fmt.Errorf("the input has to be an integer")
	}

	parsedIdx := idx - 1

	if !utils.CheckIndexExistsInSlice(len(*todos), int(parsedIdx)) {
		return fmt.Errorf("todo index is not found")
	}

	if (*todos)[parsedIdx].Completed {
		(*todos)[parsedIdx] = interfaces.Todo{
			Name:      (*todos)[parsedIdx].Name,
			Completed: false,
		}

		fmt.Printf("%s already toggled as uncompleted ✔️\n", (*todos)[parsedIdx].Name)
	} else {
		(*todos)[parsedIdx] = interfaces.Todo{
			Name:      (*todos)[parsedIdx].Name,
			Completed: true,
		}

		fmt.Printf("%s already toggled as completed ✔️\n", (*todos)[parsedIdx].Name)
	}

	return nil
}

func Search(todos *[]interfaces.Todo, todoName string) error {
	if todoName == "" {
		return fmt.Errorf("todo cannot be empty")
	}

	var foundTodos []struct {
		index int
		name  string
	}

	for i, todo := range *todos {
		if strings.Contains(todo.Name, todoName) {
			foundTodos = append(foundTodos, struct {
				index int
				name  string
			}{
				index: i + 1,
				name:  todo.Name,
			})
		}
	}

	if len(foundTodos) < 1 {
		fmt.Println("no result found")
		return nil
	}

	fmt.Printf("found %v result(s) matched\n", len(foundTodos))

	for _, foundTodo := range foundTodos {
		fmt.Printf("%v) %s\n", foundTodo.index, foundTodo.name)
	}

	return nil
}
