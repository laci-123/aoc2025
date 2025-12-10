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
	splitters map[Splitter]bool
	start Splitter
	height int
}

func parseInput(input []byte) Diagram {
	var start Splitter
	splitters := make(map[Splitter]bool)
	x := 0
	y := 0
	for _, c := range input {
		switch c {
			case 'S':
				start.x = x
				start.y = y
			case '^':
				splitters[Splitter{x, y}] = false
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

func beam(x, y, height int, splitters map[Splitter]bool) {
	fmt.Println(x, y, height)
	for y < height {
		sp := Splitter{x, y}
		if _, ok := splitters[sp]; ok {
			splitters[sp] = true
			beam(x-1, y+1, height, splitters)
			beam(x+1, y+1, height, splitters)
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
		if v {
			s += 1
		}
	}
	return s
}

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	diagram := parseInput(input)
	result1 := part1(diagram)
	fmt.Println("part1 result: ", result1)
}
