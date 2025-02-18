package interaction

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/J-yriz/todo-list-cli/cmd/utility"
	"github.com/J-yriz/todo-list-cli/model"
)

func CreateTodo() {

	resultInput, err := utility.CmdInput("Masukkan Todo [Note : Jika ingin membuat new line bisa mengetikan \\n]")
	if err != nil {
		fmt.Println(resultInput)
		panic(err.Error())
	}

	option := []string{"Simpan", "Tidak Simpan"}
	resultSelect, err := utility.CmdSelect("Konfirmasi", option)
	if err != nil {
		panic(err.Error())
	}

	data := GetUsersTodo()

	switch resultSelect {
	case option[0]:
		data.DataTodo = append(data.DataTodo, model.DataTodo{ID: int16(len(data.DataTodo) + 1), Feild: resultInput, CreatedAt: time.Now().Format("2006-01-02 15:04:05")})
		newData, _ := json.MarshalIndent(data, "", " ")

		file, _ := os.Create("data/users.json")
		if _, err = fmt.Fprint(file, string(newData)); err != nil {
			panic(err.Error())
		}

		fmt.Print("Todo berhasil disimpan\n\n")
	case option[1]:
		fmt.Print("Todo tidak disimpan\n\n")
	}

}
