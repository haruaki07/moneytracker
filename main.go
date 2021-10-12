package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/google/shlex"
)

func getCommand() string {
	fmt.Print("\033[36m", "moneytrack> ", "\033[0m")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	return input
}

func printError(msg ...string) {
	if len(msg) > 0 && msg[0] != "" {
		for _, m := range msg {
			fmt.Println(m)
		}
	} else {
		fmt.Println("Invalid input command!")
	}
}

func main() {
	var records = GetAll()

	input := getCommand()

	for input != "exit" {
		args, err := shlex.Split(input)
		if err != nil {
			printError()
			input = getCommand()
		}

		if len(args) > 0 {
			switch args[0] {
			case "add":
				params := args[1:]
				success, msg := records.AddRecord(params)
				if success == false {
					printError(msg)
				}
			case "get":
				if len(args) > 1 {
					ids := args[1:]
					fmt.Printf("%s", records.GetById(ids))
				} else {
					fmt.Printf("%s", records)
				}
				break
			case "clear":
				records.Clear()
			default:
				printError()
			}
		}

		input = getCommand()
	}

	os.Exit(0)
}
