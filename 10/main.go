package main

import (
	"aoc_2022/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type instruction interface {
	execute(x int) (cycles int, xVal int)
}

type addx struct {
	amount int
}

func (a addx) execute(x int) (cycles int, xVal int) {
	return 2, x + a.amount
}

type noop struct {
}

func (n noop) execute(x int) (cycles int, xVal int) {
	return 1, x
}

type cpu struct {
	x            int
	cycles       int
	instructions []instruction
	singnals     []int
	crtPos       int
	currentLine  []rune
}

func (c *cpu) addToLine() {
	if c.x-1 == c.crtPos {
		c.currentLine[c.x-1] = '#'
	}

	if c.x+1 == c.crtPos {
		c.currentLine[c.x+1] = '#'
	}

	if c.x == c.crtPos {
		c.currentLine[c.x] = '#'
	}

}

func (c *cpu) runInstruction(instructionIndex int) {
	cycles, x := c.instructions[instructionIndex].execute(c.x)

	for i := 0; i < cycles; i++ {
		c.cycles++

		if ((c.cycles - 20) % 40) == 0 {
			c.singnals = append(c.singnals, c.x*c.cycles)
		}

		if (c.cycles % 40) == 0 {
			c.addToLine()
			fmt.Println(string(c.currentLine))
			c.currentLine = createLine()
			c.crtPos = 0
		} else {
			c.addToLine()

			c.crtPos++
		}

		if i == cycles-1 {
			c.x = x
		}

		// fmt.Println("Cycle: ", c.cycles, "x value: ", c.x, "Instruction number: ", instructionIndex)
	}

}

func newCPU() cpu {
	emptyLine := createLine()
	return cpu{
		x:            1,
		cycles:       0,
		instructions: []instruction{},
		singnals:     []int{},
		crtPos:       0,
		currentLine:  emptyLine,
	}
}

func main() {
	start := time.Now()

	simulateCPU()

	fmt.Println("Binomial took: ", time.Since(start))
}

func getInstructions() cpu {
	scanner := utils.ReadFileAsScanner("input")

	cpu := newCPU()

	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		if instruction[0] == "noop" {
			newNoop := noop{}
			cpu.instructions = append(cpu.instructions, newNoop)
		}

		if instruction[0] == "addx" {
			val, _ := strconv.Atoi(instruction[1])
			newAddx := addx{
				amount: val,
			}

			cpu.instructions = append(cpu.instructions, newAddx)
		}
	}

	return cpu
}

func simulateCPU() {
	cpu := getInstructions()

	for i := 0; i < len(cpu.instructions); i++ {
		cpu.runInstruction(i)

	}

	sum := 0
	for i := 0; i < len(cpu.singnals); i++ {
		sum += cpu.singnals[i]
	}
	fmt.Println(cpu.singnals)
	fmt.Println(sum)
}

func createLine() []rune {
	line := make([]rune, 40)
	for i := 0; i < 40; i++ {
		line[i] = '.'
	}

	return line
}
