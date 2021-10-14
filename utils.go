package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func RandID(n int) string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func ValidateDate(str string) bool {
	dateRegex := regexp.MustCompile(`\d{2}/\d{2}/\d{4} \d{2}:\d{2}`)
	return dateRegex.MatchString(str)
}

func GetFormattedDate(t time.Time) string {
	str := fmt.Sprintf("%d/%d/%d %d:%02d",
		t.Day(),
		t.Month(),
		t.Year(),
		t.Hour(),
		t.Minute(),
	)

	return str
}

func PrintError(msg ...string) {
	if len(msg) > 0 && msg[0] != "" {
		for _, m := range msg {
			fmt.Println(m)
		}
	} else {
		fmt.Println("Invalid input command!")
	}
}

func InArray(str string, arr []string) bool {
	var found = false

	for _, v := range arr {
		if v == str {
			found = true
		}
	}

	return found
}
