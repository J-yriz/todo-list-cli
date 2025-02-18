package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/J-yriz/todo-list-cli/cmd/utility"
)

func Welcome() {

	// Membaca file data/users.json
	file, err := os.Open("data/users.json")
	if err != nil {
		fmt.Println("-============================-")
		fmt.Println("|  Welcome to Todo List CLI  |")
		fmt.Println("-============================-")

		// Jika file tidak ditemukan, maka akan meminta input username
		InputUser()
		return
	}

	// Menutup file
	defer utility.CloseFile(file)
	content, err := os.ReadFile(file.Name())
	if err != nil {
		panic(err.Error())
	}

	data := map[string]any{}
	err = json.Unmarshal(content, &data)
	if err != nil {
		panic(err.Error())
	}

	if value, exists := data["Username"]; exists {
		fmt.Println("-============================-")
		fmt.Println("|  Welcome to Todo List CLI  |")
		fmt.Println("-============================-")
		fmt.Printf("Username: %s  \n\n", value)
	}

}
