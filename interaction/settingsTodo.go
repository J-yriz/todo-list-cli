package interaction

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/J-yriz/todo-list-cli/cmd/utility"
)

func SettingsTodo() {

	option := []string{"Change Username", "Back"}
	resultSelect, err := utility.CmdSelect("Select Settings", option)
	if err != nil {
		panic(err.Error())
	}

	switch resultSelect {
	case option[0]:
		data := GetUsersTodo()
		resultEdit, _ := utility.CmdEdit("Change Username", data.Username)
		data.Username = resultEdit
		newData, _ := json.MarshalIndent(data, "", " ")
		file, _ := os.Create("data/users.json")
		fmt.Fprint(file, string(newData))

		fmt.Print("Change Username\n\n")
	}
}
