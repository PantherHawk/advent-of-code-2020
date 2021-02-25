package main

func validatePasswords(passwords []password) int {
	// initialize count of valid passwords
	validPasswordCount := 0
	// for each password, check if there are targetCount target letter
	for i := range passwords {
		pswd := passwords[i]
		targetCount := 0
		for _, char := range pswd.password {
			if string(char) == pswd.searchChar {
				targetCount++
			}
		}
		if targetCount >= pswd.min && targetCount <= pswd.max {
			validPasswordCount++
		}
	}
	// if in bounds, increment count
	return validPasswordCount
}
