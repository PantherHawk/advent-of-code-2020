package main

import "testing"

var input = []string{
	"1-3 a: abcde",
	"1-3 b: cdefg",
	"2-9 c: ccccccccc"}

func TestPasswordPhilosophy(t *testing.T) {
	got := validatePasswords(parsePasswords(input))
	if got != 2 {
		t.Errorf("Valid Passwords = %d; want 2", got)
	}
}
