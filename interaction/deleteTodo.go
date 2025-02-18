package interaction

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/J-yriz/todo-list-cli/cmd/utility"
)

func DeleteTodo() {

	data := GetUsersTodo()

	dataSelect := make([]string, 0, len(data.DataTodo))
	for i, v := range data.DataTodo {
		if len(v.Feild) > 50 {
			v.Feild = v.Feild[:50] + "..."
		}
		dataSelect = append(dataSelect, fmt.Sprintf("%d. %s | %s", i+1, v.Feild, v.CreatedAt))
	}

	resultSelect, err := utility.CmdSelect("Select Todo to delete", dataSelect)
	if err != nil {
		panic(err.Error())
	}

	numberInt, _ := strconv.Atoi(strings.Split(resultSelect, ".")[0])
	data.DataTodo = append(data.DataTodo[:numberInt-1], data.DataTodo[numberInt:]...)
	newData, _ := json.MarshalIndent(data, "", " ")
	files, _ := os.Create("data/users.json")
	fmt.Fprint(files, string(newData))

	fmt.Print("Todo berhasil dihapus\n\n")

}
