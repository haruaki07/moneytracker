package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Records struct {
	Records []Record `json:"records"`
}

type Record struct {
	Id          string `json:"id"`
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	Type        string `json:"type"`
	At          string `json:"at"`
}

func (r Records) String() string {
	var sb strings.Builder

	for i, v := range r.Records {
		str := fmt.Sprintf("%d. %s\n", i+1, v.String())
		sb.WriteString(str)
	}

	return sb.String()
}

func (r Record) String() string {
	var sb strings.Builder

	str := fmt.Sprintf("[%s] %s | %d | %s ", r.Id, r.At, r.Amount, r.Type)
	sb.WriteString(str)

	if r.Description != "" {
		sb.WriteString(fmt.Sprintf("| %s ", r.Description))
	}

	return sb.String()
}

func GetAll() Records {
	var data []byte
	var records Records

	jsonFile, err := os.Open("data.json")
	if os.IsNotExist(err) {
		data = []byte(`{ "users": [] }`)
	} else if err == nil {
		byteValue, _ := ioutil.ReadAll(jsonFile)
		data = byteValue
	}

	json.Unmarshal(data, &records)

	defer jsonFile.Close()

	return records
}

func (r *Records) GetById(ids []string) *Records {
	var results Records
	for _, record := range r.Records {
		for _, id := range ids {
			if record.Id == id {
				results.Records = append(results.Records, record)
			}
		}
	}

	return &results
}

func (r *Records) Clear() {
	var response string

	fmt.Print("This can't be undone! Are you sure (Y/N)? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	response = scanner.Text()

	if strings.ToLower(response) == "y" {
		r.Records = nil
		r.SaveToFile()
	}
}

func (r *Records) AddRecord(params []string) (bool, string) {
	formatErrorMsg := "Invalid input format!"

	if len(params) <= 1 {
		return false, formatErrorMsg
	}

	id := RandID(5)
	var (
		amount   int64
		desc, at string
	)
	_type := "in"

	for i, v := range params {
		if v == "-a" {
			if len(params) <= i+1 {
				return false, formatErrorMsg
			}

			amount, _ = strconv.ParseInt(params[i+1], 10, 64)
		}

		if v == "-d" {
			if len(params) <= i+1 {
				return false, formatErrorMsg
			}

			desc = params[i+1]
		}

		if v == "--out" {
			_type = "out"
		}

		if v == "-t" {
			if len(params) <= i+1 {
				return false, ""
			}

			if ValidateDate(params[i+1]) == false {
				return false, "Invalid date format! \nValid format: \"dd/mm/yyyy hh:MM\" (with double quotes)"
			}

			at = params[i+1]
		}
	}

	now := time.Now()
	at = fmt.Sprintf("%d/%d/%d %d:%d",
		now.Day(),
		now.Month(),
		now.Year(),
		now.Hour(),
		now.Minute(),
	)

	record := Record{
		Id:          id,
		Amount:      amount,
		Description: desc,
		Type:        _type,
		At:          at,
	}

	r.Records = append(r.Records, record)

	fmt.Printf("Added new record!\n%s\n", record)

	r.SaveToFile()

	return true, ""
}

func (r *Records) SaveToFile() {
	data, _ := json.Marshal(r)

	ioutil.WriteFile("data.json", data, 0644)
}
