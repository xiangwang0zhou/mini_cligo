package commands

import (
	"fmt"
)

func CallbackHelp() {
	fmt.Println("These is minicli's help menu")
	avavilableCommand := GetCommands()
	for _, cmd := range avavilableCommand {
		fmt.Printf("-%s:%s\n", cmd.name, cmd.description)
	}
}
