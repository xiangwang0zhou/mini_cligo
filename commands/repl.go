package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Repl() {
	Scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">")
		Scanner.Scan()
		content := Scanner.Text()
		cleared := ClearInput(content)
		if len(cleared) == 0 {
			continue
		} //get command
		commandName := cleared[0]
		availableCommands := GetCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invaild command")
			continue
		}
		command.callback()
	}
}

func ClearInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type CliCommand struct {
	name        string
	description string
	callback    func()
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "get all commands and description",
			callback:    CallbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "stdander way to exit",
			callback:    CallbackExit,
		},
	}
}
