package main

import (
	"fmt"
	"io"
	"os"
)

type Splitter struct {
	x int
	y int
}

type Diagram struct {
	splitters map[Splitter]int
	start Splitter
	height int
}

func parseInput(input []byte) Diagram {
	var start Splitter
	splitters := make(map[Splitter]int)
	x := 0
	y := 0
	for _, c := range input {
		switch c {
			case 'S':
				start.x = x
				start.y = y
			case '^':
				splitters[Splitter{x, y}] = 0
			case '\n':
				x = -1
				y += 1
			default:
				// do nothing
		}
		x += 1
	}
	return Diagram{splitters: splitters, start: start, height: y}
}

func beam(x, y, height int, splitters map[Splitter]int) {
	for y < height {
		sp := Splitter{x, y}
		if v, ok := splitters[sp]; ok {
			if v == 0 {
				splitters[sp] = 1
				beam(x-1, y+1, height, splitters)
				beam(x+1, y+1, height, splitters)
			}
			break
		} else {
			y += 1
		}
	}
}

func part1(diagram Diagram) int {
	beam(diagram.start.x, diagram.start.y, diagram.height, diagram.splitters)
	s := 0
	for _, v := range diagram.splitters {
		if v == 1 {
			s += 1
		}
	}
	return s
}

func beam2(x, y, height int, splitters map[Splitter]int) int {
	for y < height {
		sp := Splitter{x, y}
		if v, ok := splitters[sp]; ok {
			count := splitters[sp]
			if v == 0 {
				count += beam2(x-1, y+1, height, splitters)
				count += beam2(x+1, y+1, height, splitters)
				splitters[sp] = count
			}
			return count
		} else {
			y += 1
		}
	}
	return 1
}

func part2(diagram Diagram) int {
	return beam2(diagram.start.x, diagram.start.y, diagram.height, diagram.splitters)
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	diagram1 := parseInput(input)
	result1 := part1(diagram1)
	fmt.Println("part1 result: ", result1)

	diagram2 := parseInput(input)
	result2 := part2(diagram2)
	fmt.Println("part2 result: ", result2)
}
