package main

import (
	"aoc_2022/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	findOverlappingShifts()

	fmt.Println("Binomial took: ", time.Since(start))
}

func findOverlappingShifts() {
	scanner := utils.ReadFileAsScanner("input")

	sumOfDuplicatedShifts := 0
	sumOfOverlappingShifts := 0

	for scanner.Scan() {
		shift := splitShifts(scanner.Text())

		if (shift[0][0] <= shift[1][0] && shift[0][1] >= shift[1][1]) || (shift[1][0] <= shift[0][0] && shift[1][1] >= shift[0][1]) {
			fmt.Println(shift)
			sumOfDuplicatedShifts++
		}

		if !(shift[0][1] < shift[1][0] || shift[1][1] < shift[0][0]) {
			sumOfOverlappingShifts++
		}
	}

	fmt.Println("Duplicated shifts: ", sumOfDuplicatedShifts)
	fmt.Println("Overlapping shifts: ", sumOfOverlappingShifts)
}

func splitShifts(shiftData string) [][]int {
	shifts := strings.Split(shiftData, ",")

	shiftNumbers := make([][]int, 2)

	for i := 0; i < len(shifts); i++ {
		partnerData := strings.Split(shifts[i], "-")
		shiftNumbers[i] = make([]int, 2)
		for j := 0; j < len(partnerData); j++ {
			val, _ := strconv.Atoi(partnerData[j])
			shiftNumbers[i][j] = val
		}

	}

	return shiftNumbers
}
