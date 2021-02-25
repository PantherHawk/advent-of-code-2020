package main

import (
	"fmt"
	"regexp"
)

type password struct {
	password   string
	min        int
	max        int
	searchChar string
}

var regex = regexp.MustCompile(`(^\d+)-(\d+) (\w): (\w+)`)

func main() {
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc"}
	passwords := parsePasswords(input)
	result := validatePasswords(passwords)
	fmt.Print("RESULT =>", result)
}
