package main

import (
	"aoc_2022/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x int
	y int
}

func (p *point) toString() string {
	return strconv.Itoa(p.x) + " " + strconv.Itoa(p.y)
}

type state struct {
	length int
	knots  []point
	path   map[string]point
}

func (s *state) Move(dir string) {
	switch dir {
	case "U":
		s.knots[0].y++
	case "D":
		s.knots[0].y--
	case "L":
		s.knots[0].x--
	case "R":
		s.knots[0].x++
	}
}

func (s *state) MoveTail() {
	for i := 1; i < s.length; i++ {
		delta := point{s.knots[i-1].x - s.knots[i].x, s.knots[i-1].y - s.knots[i].y}
		if math.Abs(float64(delta.x)) <= 1 && math.Abs(float64(delta.y)) <= 1 {
			return
		}

		if delta.y > 0 {
			s.knots[i].y++
		} else if delta.y < 0 {
			s.knots[i].y--
		}
		if delta.x > 0 {
			s.knots[i].x++
		} else if delta.x < 0 {
			s.knots[i].x--
		}
	}
	s.AddPoint(s.knots[s.length-1])
}

func (s *state) AddPoint(p point) {
	fmt.Println("adding point to set")
	s.path[p.toString()] = p
}

func main() {
	start := time.Now()

	findUniqueLocations(10)

	fmt.Println("Binomial took: ", time.Since(start))
}

func findUniqueLocations(amountOfKnots int) {
	scanner := utils.ReadFileAsScanner("input")

	knots := make([]point, amountOfKnots)

	for i := 0; i < len(knots); i++ {
		newKnot := point{
			x: 0,
			y: 0,
		}
		knots[i] = newKnot
	}

	ropeState := state{
		length: amountOfKnots,
		knots:  knots,
		path:   make(map[string]point),
	}

	ropeState.AddPoint(ropeState.knots[ropeState.length-1])

	for scanner.Scan() {
		command := strings.Split(scanner.Text(), " ")
		direction := command[0]
		steps, _ := strconv.Atoi(command[1])

		for i := 0; i < steps; i++ {

			ropeState.Move(direction)
			ropeState.MoveTail()

		}

	}

	fmt.Println(len(ropeState.path))
}
