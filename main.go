/**
* .---------------------------------------------------------------------------.
* |  _     _____    _    ____  _   _    ____  ___  _        _    _   _  ____  |
* | | |   | ____|  / \  |  _ \| \ | |  / ___|/ _ \| |      / \  | \ | |/ ___| |
* | | |   |  _|   / _ \ | |_) |  \| | | |  _| | | | |     / _ \ |  \| | |  _  |
* | | |___| |___ / ___ \|  _ <| |\  | | |_| | |_| | |___ / ___ \| |\  | |_| | |
* | |_____|_____/_/   \_\_| \_\_| \_|  \____|\___/|_____/_/   \_\_| \_|\____| |
* '---------------------------------------------------------------------------'
 */

package main

import (
	"github.com/J-yriz/todo-list-cli/cmd"
	"github.com/J-yriz/todo-list-cli/interaction"
)

func main() {

	cmd.Welcome()

loop:
	for {
		switch cmd.Navigaiton() {
		case cmd.Items[0]:
			interaction.CreateTodo()
		case cmd.Items[1]:
			interaction.EditTodo()
		case cmd.Items[2]:
			interaction.DeleteTodo()
		case cmd.Items[3]:
			interaction.ViewTodo()
		case cmd.Items[4]:
			interaction.SettingsTodo()
		case cmd.Items[5]:
			break loop
		}
	}

}
