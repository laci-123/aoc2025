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

// modulo function: unllike a%b (remainder) the result is always non-negative
func mod(a int, b int) int {
    m := a % b
    if m < 0 {
        m += b
    }
    return m
}

func part1() (int, error) {
	dial := 50
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		rotation, err := parseRotation(text)
		if err != nil {
			return 0, err
		}
		dial = mod(dial + rotation, 100)
		if dial == 0 {
			sum += 1
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading stdin")
	}
	return sum, nil
}

func main() {
	result, err := part1()
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Println(result)
	}
}
