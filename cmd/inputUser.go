package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/J-yriz/todo-list-cli/cmd/utility"
)

type dataSave struct {
	Username string
}

func InputUser() {

	if _, err := os.Stat("data"); os.IsNotExist(err) {
		if err := os.Mkdir("data", os.ModePerm); err != nil {
			panic(err.Error())
		}

		result, err := utility.CmdInput("Masukkan Username")
		if err != nil {
			panic(err.Error())
		}

		file, _ := os.Create("data/users.json")
		jsonData, err := json.MarshalIndent(dataSave{result}, "", " ")
		if err != nil {
			panic(err.Error())
		}

		if _, err = fmt.Fprint(file, string(jsonData)); err != nil {
			panic(err.Error())
		}
	}

}
