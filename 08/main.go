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

	findVisibleTrees()
	findBestSenicScore()

	fmt.Println("Binomial took: ", time.Since(start))
}

func parseInput() [][]int {
	scanner := strings.Split(utils.ReadFileAsString("input"), "\n")

	forest := make([][]int, len(scanner))

	for i := 0; i < len(scanner); i++ {
		forest[i] = make([]int, len(scanner[i]))

		line := strings.Split(scanner[i], "")

		for j := 0; j < len(line); j++ {
			heigth, _ := strconv.Atoi(line[j])
			forest[i][j] = heigth
		}
	}

	return forest
}

func findVisibleTrees() {
	forest := parseInput()

	ableToSee := 0
	for s := 0; s < len(forest); s++ {
		for a := 0; a < len(forest); a++ {
			selectedTree := forest[s][a]
			if s == 0 || a == 0 || s == len(forest)-1 || a == len(forest)-1 {
				ableToSee = ableToSee + 1
			} else if selectedTree > 0 {
				if utils.FindMaxInArray(forest[s][a+1:len(forest)]) < selectedTree || utils.FindMaxInArray(forest[s][0:a]) < selectedTree {
					ableToSee++
					continue
				}

				verticalRow := utils.GetVerticalArray(forest, a)

				if utils.FindMaxInArray(verticalRow[0:s]) < selectedTree || utils.FindMaxInArray(verticalRow[s+1:len(forest)]) < selectedTree {
					ableToSee++
					continue
				}

			}

		}
	}
	fmt.Println("Visible trees: ", ableToSee)
}

func findBestSenicScore() {
	forest := parseInput()

	higestSenicScore := 0
	for s := 0; s < len(forest); s++ {
		for a := 0; a < len(forest); a++ {
			selectedTree := forest[s][a]

			verticalRow := utils.GetVerticalArray(forest, a)

			senicScore := calculateSenicScoreTo(forest[s][0:a], selectedTree) * calculateSenicScoreFrom(forest[s][a+1:len(forest)], selectedTree) * calculateSenicScoreTo(verticalRow[0:s], selectedTree) * calculateSenicScoreFrom(verticalRow[s+1:len(forest)], selectedTree)

			if senicScore > higestSenicScore {

				higestSenicScore = senicScore
			}

		}
	}
	fmt.Println("Visible trees: ", higestSenicScore)
}

func calculateSenicScoreTo(row []int, height int) int {
	if len(row) == 0 {
		return 0
	}

	counter := 0
	for i := len(row) - 1; i > 0; i-- {
		counter++
		if row[i] >= height {
			return counter
		}
	}

	return len(row)
}

func calculateSenicScoreFrom(row []int, height int) int {
	if len(row) == 0 {
		return 0
	}

	for i := 0; i < len(row); i++ {
		if row[i] >= height {
			return i + 1
		}
	}

	return len(row)
}
