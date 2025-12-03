package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type BatteryBank struct {
	batteries []int
}

func readInput() ([]BatteryBank, error) {
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	bank := make([]BatteryBank, 0)
	for line := range strings.Lines(string(content)) {
		trimmedLine := strings.Trim(line, " \n")
		batteries := make([]int, 0)
		for i := 0; i < len(trimmedLine); i += 1 {
			x, err := strconv.Atoi(trimmedLine[i:(i+1)])
			if err != nil {
				return nil, err
			}
			batteries = append(batteries, x)
		}
		bank = append(bank, BatteryBank{batteries})
	}
	return bank, nil
}

func part1(banks []BatteryBank) int {
	sum := 0
	for _, bank := range banks {
		max := 0
		for i := 0; i < len(bank.batteries); i += 1 {
			for j := i + 1; j < len(bank.batteries); j += 1 {
				x := bank.batteries[i] * 10 + bank.batteries[j]
				if x > max {
					max = x
				}
			}
		}			
		sum += max
	}
	return sum
}


func main() {
	banks, err := readInput()
	if err != nil {
		panic(err)
	} 

	result1 := part1(banks)
	fmt.Println("part 1 result: ", result1)

	// result2 := part2(ranges)
	// fmt.Println("part 2 result: ", result2)
}
