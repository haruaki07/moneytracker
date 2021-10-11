package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"strconv"
	"github.com/google/shlex"
)

type Record struct {
	Id 			string
	Amount		int64
	Description	string
	Type		string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

func (r Record) String() string {
	return fmt.Sprintf("[%s] %d (%s) %s", r.Id, r.Amount, r.Type, r.Description)
}

func addRecord(params []string) {
	id := RandID(5)
	var amount int64
	var desc string
	_type := "in"
	createdAt, updatedAt := time.Now(), time.Now()
	
	for i, v := range params {
		if v == "-a" {
			amount, _ = strconv.ParseInt(params[i + 1], 10, 64)
		}
		
		if v == "-d" {
			desc = params[i + 1]
		}
		
		if v == "--out" {
			_type = "out"
		}
	}
	
	record := Record{
		Id: id,
		Amount: amount,
		Description: desc,
		Type: _type,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	
	fmt.Println(record)
}

func processCommand(args []string) {
	cmd := args[0]

	switch cmd {
	case "exit":
		os.Exit(0)
	case "add":
		params := args[1:]
		addRecord(params)
	}
}

func getCommand() {
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	input := scanner.Text()

	args, err := shlex.Split(input)
	if err != nil {
		os.Exit(2)
	}
	
	processCommand(args)	
}

func main() {
	getCommand()
}
