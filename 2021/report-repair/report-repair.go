package main

import (
	"fmt"
	"sort"
)

func main() {

	years := []int{1721, 979, 366, 299, 675, 1456}
	repairReportOptimized(years, 2020)
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
	fmt.Println(product)
	return product
}

func repairReportOptimized(yearList []int, target int) int {
	sort.Ints(yearList)
	var product int
	for i := 0; i < len(yearList); i++ {
		for j := len(yearList) - 1; j > i; j-- {
			x := yearList[i]
			y := yearList[j]
			sum := x + y
			if sum == target {
				product = x * y
			}
		}
	}
	fmt.Println(product)
	return product
}
