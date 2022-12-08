package main

import (
	"aoc_2022/utils"
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	sortShip()

	fmt.Println("Binomial took: ", time.Since(start))
}

type ship struct {
	containers map[int]*utils.StringStack
}

func sortShip() {
	input := utils.SplitByEmptyNewline(utils.ReadFileAsString("input.txt"))
	ship := ship{containers: parseShip(input[0])}

	scanner := bufio.NewScanner(strings.NewReader(input[1]))

	for scanner.Scan() {
		command := getMoveCommand(scanner.Text())
		ship.moveCrates9001(command[0], command[1], command[2])
	}

	for i := 1; i <= 9; i++ {
		x, _ := ship.containers[i].Pop()
		fmt.Print(x)
	}
	fmt.Println(" ")

}

func (s *ship) moveCrates(amount int, from int, to int) {
	for i := 0; i < amount; i++ {
		container, err := s.containers[from].Pop()
		if err != nil {
			log.Fatal(err)
			break
		}

		s.containers[to].Push(container)
	}
}

func (s *ship) moveCrates9001(amount int, from int, to int) {
	containers := make([]string, amount)
	for i := 0; i < amount; i++ {
		container, err := s.containers[from].Pop()
		if err != nil {
			log.Fatal(err)
			break
		}
		containers[i] = container
	}

	for i := amount - 1; i >= 0; i-- {

		s.containers[to].Push(containers[i])
	}
}

func parseShip(shipData string) map[int]*utils.StringStack {
	input := strings.Split(shipData, "\n")
	ship := make(map[int]*utils.StringStack)

	res := regexp.
		MustCompile(`\d+`).
		FindAllString(input[len(input)-1], -1)

	for i := 0; i < len(res); i++ {
		x, _ := strconv.Atoi(res[i])
		ship[x] = utils.NewStringStack()
	}

	indexes := make(map[int]int)
	a := strings.Split(input[len(input)-1], "")
	for i := 0; i < len(a); i++ {
		x, err := strconv.Atoi(a[i])
		if err == nil {
			indexes[i] = x
		}
	}

	for i := len(input) - 2; i >= 0; i-- {
		split := strings.Split(input[i], "")

		for j := 0; j < len(split); j++ {
			if utils.IsLetter(split[j]) {
				ship[indexes[j]].Push(split[j])
			}
		}
	}

	return ship
}

func getMoveCommand(str string) []int {
	res := regexp.
		MustCompile(`\d+`).
		FindAllString(str, -1)

	commands := make([]int, 3)
	for i, s := range res {
		x, _ := strconv.Atoi(s)
		commands[i] = x
	}

	return commands
}
