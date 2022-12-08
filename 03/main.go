package main

import (
	"aoc_2022/utils"
	"fmt"
	"strings"
	"time"
	"unicode"
)

func main() {
	start := time.Now()

	sortCompartments()
	sortGroups()

	fmt.Println("Binomial took: ", time.Since(start))
}

var priority map[string]int = calculatePriorityMap()

func sortCompartments() {
	scanner := utils.ReadFileAsScanner("input")

	sum := 0

	for scanner.Scan() {
		compartments := strings.Split(scanner.Text(), "")
		firstCompartment := compartments[0 : len(compartments)/2]
		secondCompartment := compartments[len(compartments)/2:]

		commonGroup := make(map[string]int)
		commonGroup = compareInventories(commonGroup, firstCompartment, priority)
		commonGroup = compareInventories(commonGroup, secondCompartment, priority)

		for _, element := range commonGroup {
			sum = sum + element
		}

	}

	fmt.Println(sum)
}

func sortGroups() {
	inventorySlice := strings.Split(utils.ReadFileAsString("input"), "\n")

	sum := 0

	for elfIndex := 0; elfIndex < len(inventorySlice); elfIndex = elfIndex + 3 {

		commonGroup := make(map[string]int)

		for groupindex := 0; groupindex < 3; groupindex++ {
			trueIndex := groupindex + elfIndex
			commonGroup = compareInventories(commonGroup, strings.Split(string(inventorySlice[trueIndex]), ""), priority)
		}

		for _, element := range commonGroup {
			sum = sum + element
		}

	}
	fmt.Println(sum)
}

func compareInventories(commonGroup map[string]int, inventory []string, PriorityMap map[string]int) map[string]int {
	common := make(map[string]int)
	if len(commonGroup) == 0 {
		for inventoryIndex := 0; inventoryIndex < len(inventory); inventoryIndex++ {
			delete(commonGroup, inventory[inventoryIndex])
			common[inventory[inventoryIndex]] = PriorityMap[inventory[inventoryIndex]]
		}
	}

	for inventoryIndex := 0; inventoryIndex < len(inventory); inventoryIndex++ {
		if commonGroup[inventory[inventoryIndex]] >= 1 {
			delete(commonGroup, inventory[inventoryIndex])
			common[inventory[inventoryIndex]] = PriorityMap[inventory[inventoryIndex]]
		}
	}

	return common
}

func calculatePriorityMap() map[string]int {
	Priority := make(map[string]int)

	lowercase := 1
	uppercase := 27

	for r := 'a'; r <= 'z'; r++ {
		R := unicode.ToUpper(r)

		Priority[string(r)] = lowercase
		Priority[string(R)] = uppercase

		lowercase = lowercase + 1
		uppercase = uppercase + 1
	}

	return Priority
}
