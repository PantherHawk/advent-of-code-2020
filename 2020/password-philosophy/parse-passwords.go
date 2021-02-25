package main

import "strconv"

func parsePasswords(input []string) []password {
	passwords := make([]password, len(input))
	for i := range passwords {
		split := regex.FindStringSubmatch(input[i])
		min, _ := strconv.Atoi(split[1])
		max, _ := strconv.Atoi(split[2])
		if len(split) > 3 {
			passwords[i] = password{split[4], min, max, split[3]}
		}
	}
	return passwords
}
