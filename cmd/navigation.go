package cmd

import "github.com/J-yriz/todo-list-cli/cmd/utility"

var Items = []string{"Create Todo", "Edit Todo", "Delete Todo", "View Todo", "Settings", "Exit"}

func Navigaiton() string {

	result, err := utility.CmdSelect("Select Navigation", Items)

	if err != nil {
		panic(err.Error())
	}

	return result

}
