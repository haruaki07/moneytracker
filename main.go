package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"

	"github.com/google/shlex"
)

func getCommand() string {
	fmt.Print("\033[36m", "moneytrack> ", "\033[0m")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	return input
}

func main() {
	RegisterHelp()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.Exit(0)
	}()

	var records = GetAll()

	input := getCommand()

	for input != "exit" {
		args, err := shlex.Split(input)
		if err != nil {
			PrintError()
			input = getCommand()
		}

		if len(args) > 0 {
			switch args[0] {
			case "add":
				params := args[1:]
				records.AddRecord(params)
			case "get":
				if len(args) > 1 {
					params := args[1:]
					fmt.Printf("%s", records.GetById(params))
				} else {
					fmt.Printf("%s", records)
				}
			case "set":
				params := args[2:]
				records.Update(args[1], params)
			case "del":
				if len(args) > 1 {
					ids := args[1:]
					records.DeleteById(ids)
				} else {
					PrintError("Invalid input format!", "Usage: del [id]...")
				}
			case "clear":
				records.Clear()
			case "help":
				params := args[1:]
				ShowHelp(params)
			default:
				PrintError()
			}
		}

		input = getCommand()
	}

	os.Exit(0)
}
