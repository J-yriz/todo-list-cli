package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func InputUser() string {

	// os.Stdin digunakan untuk membaca input dari user, sedangkan bufio.NewReader digunakan untuk membaca input dari user dalam bentuk string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, ok := reader.ReadString('\n')

	if ok != nil {
		panic(ok.Error())
	} else {
		return name
	}
}
