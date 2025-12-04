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

func calcJoltage(digits map[int]int, newDigit int, newDigitPos int, k int) int {
	p := 1
	sum := 0
	for i := k-1; i >= 0; i -= 1 {
		if i == newDigitPos {
			sum += newDigit * p
			p *= 10
		} else if d, ok := digits[i]; ok {
			sum += d * p
			p *= 10
		}
	}
	return sum
}

func findMaxJoltage(bank BatteryBank, n int) (digits map[int]int, max int) {
	if n == 1 {
		digits := make(map[int]int)
		max := 0
		maxPos := 0
		for i := 0; i < len(bank.batteries); i += 1 {
			if bank.batteries[i] > max {
				max = bank.batteries[i]
				maxPos = i
			}
		}
		digits[maxPos] = max
		return digits, max
	} else {
		digits, max := findMaxJoltage(bank, n - 1)
		maxDigit := 0
		maxDigitPos := 0
		for i := 0; i < len(bank.batteries); i += 1 {
			if _, ok := digits[i]; ok {
				continue
			}
			x := calcJoltage(digits, bank.batteries[i], i, len(bank.batteries))
			if x > max {
				max = x
				maxDigitPos = i
				maxDigit = bank.batteries[i]
			}
		}
		digits[maxDigitPos] = maxDigit
		return digits, max
	}
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

func part2(banks []BatteryBank) int {
	sum := 0
	for _, bank := range banks {
		_, x := findMaxJoltage(bank, 12)
		sum += x
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

	result2 := part2(banks)
	fmt.Println("part 2 result: ", result2)
}
