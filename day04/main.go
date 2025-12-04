package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type PaperRoll struct{
	x int
	y int
}

func readInput() (map[PaperRoll]int, error) {
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	rolls := make(map[PaperRoll]int)
	y := 0
	for line := range strings.Lines(string(content)) {
		trimmedLine := strings.Trim(line, " \n")
		for x, c := range trimmedLine {
			if c == '@' {
				rolls[PaperRoll{x, y}] = 1
			}
		}
		y += 1
	}
	return rolls, nil
}



func part1(rolls map[PaperRoll]int) int {
	sum := 0
	for roll := range rolls {
		s := rolls[PaperRoll{x: roll.x - 1, y: roll.y - 1}]
		s += rolls[PaperRoll{x: roll.x - 1, y: roll.y}]
		s += rolls[PaperRoll{x: roll.x - 1, y: roll.y + 1}]
		s += rolls[PaperRoll{x: roll.x,     y: roll.y - 1}]
		s += rolls[PaperRoll{x: roll.x,     y: roll.y + 1}]
		s += rolls[PaperRoll{x: roll.x + 1, y: roll.y - 1}]
		s += rolls[PaperRoll{x: roll.x + 1, y: roll.y}]
		s += rolls[PaperRoll{x: roll.x + 1, y: roll.y + 1}]
		if s < 4 {
			sum += 1
		}
	}
	return sum
}

func main() {
	rolls, err := readInput()
	if err != nil {
		panic(err)
	} 

	result1 := part1(rolls)
	fmt.Println("part 1 result: ", result1)
}
