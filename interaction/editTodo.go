package interaction

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/J-yriz/todo-list-cli/cmd/utility"
)

func EditTodo() {

	data := GetUsersTodo()

	dataSelect := make([]string, 0, len(data.DataTodo))
	for i, v := range data.DataTodo {
		if len(v.Feild) > 50 {
			v.Feild = v.Feild[:50] + "..."
		}
		dataSelect = append(dataSelect, fmt.Sprintf("%d. %s | %s", i+1, v.Feild, v.CreatedAt))
	}

	resultSelect, err := utility.CmdSelect("Select Todo to edit", dataSelect)
	if err != nil {
		panic(err.Error())
	}

	numberInt, _ := strconv.Atoi(strings.Split(resultSelect, ".")[0])
	resultEdit, _ := utility.CmdEdit("Edit Todo "+strconv.Itoa(numberInt), data.DataTodo[numberInt-1].Feild)

	data.DataTodo[numberInt-1].Feild = resultEdit
	newData, _ := json.MarshalIndent(data, "", " ")
	file, _ := os.Create("data/users.json")
	fmt.Fprint(file, string(newData))

	fmt.Print("Todo berhasil diubah\n\n")

}
