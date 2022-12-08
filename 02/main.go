package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFileAsString(filepath string) string {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func ReadFileAsScanner(filepath string) *bufio.Scanner {
	scanner := bufio.NewScanner(strings.NewReader(ReadFileAsString(filepath)))

	return scanner
}

func main() {
	calculateScoreAlt()
}

// A = Paper	&& X = Paper	| 1
// B = Scissor	&& Y = Scissor	| 2
// C = Rock		&& Z = Rock		| 3

var outcomes = map[string]int{
	//losing
	"AZ": 3,
	"CY": 2,
	"BX": 1,

	//Ties
	"AX": 4,
	"BY": 5,
	"CZ": 6,

	//wins
	"AY": 8,
	"BZ": 9,
	"CX": 7,
}

func calculateScorePart1() {

	scanner := ReadFileAsScanner("input")

	score := 0

	for scanner.Scan() {
		game := strings.Join(strings.Fields(scanner.Text()), "")
		score = score + outcomes[game]
	}

	fmt.Println(score)

}

func calculateScorePart2() {

	scanner := ReadFileAsScanner("input")

	score := 0

	for scanner.Scan() {
		tScore := 0
		game := strings.Fields(scanner.Text())

		// Need to lose
		if game[1] == "X" {
			switch game[0] {
			case "A":
				tScore = outcomes["AZ"]
			case "B":
				tScore = outcomes["BX"]
			case "C":
				tScore = outcomes["CY"]
			}
		}

		// Need to draw
		if game[1] == "Y" {
			switch game[0] {
			case "A":
				tScore = outcomes["AX"]
			case "B":
				tScore = outcomes["BY"]
			case "C":
				tScore = outcomes["CZ"]
			}
		}

		// Need to win
		if game[1] == "Z" {
			switch game[0] {
			case "A":
				tScore = outcomes["AY"]
			case "B":
				tScore = outcomes["BZ"]
			case "C":
				tScore = outcomes["CX"]
			}
		}

		score = score + tScore
	}

	fmt.Println(score)

}

func calculateScoreAlt() {
	scanner := ReadFileAsScanner("input")

	points := map[string]int{"A": 0, "B": 1, "C": 2, "X": 0, "Y": 1, "Z": 2}

	scoreP1 := 0

	for scanner.Scan() {
		game := strings.Fields(scanner.Text())
		opp := points[game[0]]
		you := points[game[1]]

		draw := opp == you
		win := (you-opp+3)%3 == 1

		if win {
			scoreP1 = scoreP1 + 6
		}

		if draw {
			scoreP1 = scoreP1 + 3
		}

		scoreP1 = scoreP1 + you + 1

	}

	fmt.Println("part 1:", scoreP1, " ")
}
