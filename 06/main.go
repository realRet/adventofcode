package main

import (
	"aoc_2022/utils"
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	//Part 1
	findPacketMarker(4)

	//Part 2
	findPacketMarker(14)

	fmt.Println("Binomial took: ", time.Since(start))
}

func findPacketMarker(markerSize int) {
	data := strings.Split(utils.ReadFileAsString("input"), "")

	for i := 0; i < len(data); i++ {
		commands := make([]string, markerSize)

		for j := 0; j < len(commands); j++ {
			commands[j] = data[i+j]
		}

		// if isUniqueSequence(commands) {
		// 	fmt.Println(i + len(commands))
		// 	break
		// }

		if countDestinctCharacters(commands) == markerSize {
			fmt.Println(i + len(commands))
			break
		}
	}
}

// func isUniqueSequence(sequence []string) bool {
// 	for i := 0; i < len(sequence); i++ {
// 		for j := i + 1; j < len(sequence); j++ {
// 			if sequence[i] == sequence[j] {
// 				return false
// 			}
// 		}
// 	}

// 	return true
// }

func countDestinctCharacters(sequence []string) int {
	counter := make(map[string]int)
	for _, row := range sequence {
		counter[row]++
	}

	return len(counter)
}
