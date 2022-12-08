package main

import (
	"aoc_2022/utils"
	"fmt"
	"sort"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	getLargestInventory()

	elapsed := time.Since(start)
	fmt.Println("Binomial took: ", elapsed.Nanoseconds())

}

func getLargestInventory() {
	scanner := utils.ReadFileAsScanner("input")

	index := 0
	elfInventory := []int{0}
	higest := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			if elfInventory[index] > higest {
				higest = elfInventory[index]
			}

			elfInventory = append(elfInventory, 0)
			index = index + 1

		} else {

			number, _ := strconv.Atoi(scanner.Text())
			elfInventory[index] = elfInventory[index] + number
		}
	}

	fmt.Println("Higest: ", higest)

	sort.Ints(elfInventory)

	aLen := len(elfInventory) - 1
	highThree := elfInventory[aLen] + elfInventory[aLen-1] + elfInventory[aLen-2]

	fmt.Println("Higest Three: ", highThree)

}
