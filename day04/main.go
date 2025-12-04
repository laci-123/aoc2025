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

func removeRolls(rolls map[PaperRoll]int) int {
	toBeRemoved := make([]PaperRoll, 0)
	for roll := range rolls {
		if rolls[roll] == 0 {
			continue
		}
		s := rolls[PaperRoll{x: roll.x - 1, y: roll.y - 1}]
		s += rolls[PaperRoll{x: roll.x - 1, y: roll.y}]
		s += rolls[PaperRoll{x: roll.x - 1, y: roll.y + 1}]
		s += rolls[PaperRoll{x: roll.x,     y: roll.y - 1}]
		s += rolls[PaperRoll{x: roll.x,     y: roll.y + 1}]
		s += rolls[PaperRoll{x: roll.x + 1, y: roll.y - 1}]
		s += rolls[PaperRoll{x: roll.x + 1, y: roll.y}]
		s += rolls[PaperRoll{x: roll.x + 1, y: roll.y + 1}]
		if s < 4 {
			toBeRemoved = append(toBeRemoved, roll)
		}
	}
	for _, r := range toBeRemoved {
		rolls[r] = 0
	}
	return len(toBeRemoved)
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

func part2(rolls map[PaperRoll]int) int {
	sum := 0
	removed := -1
	for removed != 0 {
		removed = removeRolls(rolls)
		sum += removed
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

	result2 := part2(rolls)
	fmt.Println("part 2 result: ", result2)
}
