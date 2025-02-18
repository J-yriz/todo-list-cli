package interaction

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/J-yriz/todo-list-cli/model"
)

func GetUsersTodo() *model.DataAll {

	dataTodo := new(model.DataAll)
	file, err := os.ReadFile("data/users.json")
	if err != nil {
		panic(err.Error())
	}

	_ = json.Unmarshal(file, &dataTodo)

	return dataTodo

}

func ViewTodo() {

	data := GetUsersTodo()

	if len(data.DataTodo) == 0 {
		fmt.Print("Data Todo masih kosong\n\n")
		return
	}

	for i, v := range data.DataTodo {
		fmt.Println(i+1, v.Feild, v.CreatedAt)
	}
	fmt.Print("\n")

}
