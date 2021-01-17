package main

func main() {

	years := []int{1721, 979, 366, 299, 675, 1456}
	repairReport(years, 2020)
}

func repairReport(yearList []int, target int) int {
	var product int
	for _, x := range yearList {
		for _, y := range yearList {
			if x+y == target {
				product = x * y
			}
		}
	}
	return product
}
