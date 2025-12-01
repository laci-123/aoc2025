package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseRotation(rotation string) (int, error) {
	direction := rotation[0]
	var sign int
	switch direction {
		case 'L':
			sign = -1
		case 'R':
			sign = +1
		default:
			return 0, fmt.Errorf("wrong input format: expected 'L' or 'R', found %v", direction)
	}
	size, err := strconv.Atoi(rotation[1:])
	if err != nil {
		return 0, fmt.Errorf("wrong input format: %v", err)
	}
	return sign * size, nil
}

func readInput() ([]int, error) {
	rotations := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		rotation, err := parseRotation(text)
		if err != nil {
			return nil, err
		}
		rotations = append(rotations, rotation)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading stdin")
	}
	return rotations, nil
}

// modulo function: unllike a%b (remainder) the result is always non-negative
func mod(a int, b int) int {
    m := a % b
    if m < 0 {
        m += b
    }
    return m
}

func rotateDial(dial int, rotation int) (int, int) {
	sum := 0
	var step int
	if rotation < 0 {
		step = -1
	} else {
		step = 1
	}
	for i := 0; i != rotation; i += step {
		if mod(dial + i, 100) == 0 {
			sum += 1
		}
	}
	return mod(dial + rotation, 100), sum
}

func part1(rotations []int) (int, error) {
	dial := 50
	sum := 0
	for _, rotation := range rotations {
		dial = mod(dial + rotation, 100)
		if dial == 0 {
			sum += 1
		}
	}
	return sum, nil
}

func part2(rotations []int) (int, error) {
	dial := 50
	sum := 0
	for _, rotation := range rotations {
		var s int
		dial, s = rotateDial(dial, rotation)
		sum += s
	}
	return sum, nil
}

func main() {
	rotations, err := readInput()
	if err != nil {
		panic(err)
	} 

	result1, err := part1(rotations)
	if err != nil {
		panic(err)
	} 
	fmt.Println("part 1 result: ", result1)

	result2, err := part2(rotations)
	if err != nil {
		panic(err)
	}
	fmt.Println("part 2 result: ", result2)
}
