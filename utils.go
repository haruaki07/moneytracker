package main

import (
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
