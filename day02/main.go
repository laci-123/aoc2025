package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	from int
	to int
}

func readInput() ([]Range, error) {
	line, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	ranges := make([]Range, 0)
	for r := range strings.SplitSeq(string(line), ",") {
		from_to := strings.Split(r, "-")
		from, err := strconv.Atoi(strings.Trim(from_to[0], " \n"))
		if err != nil {
			return nil, err
		}
		to, err := strconv.Atoi(strings.Trim(from_to[1], " \n"))
		if err != nil {
			return nil, err
		}
		ranges = append(ranges, Range{from: from, to: to})
	}
	return ranges, nil
}

func isRepeated(s string) bool {
	m := len(s)/2
	return s[:m] == s[m:]
}

func isRepeated2(s string) bool {
	outer:
	for i := 1; i < len(s); i += 1 {
		t := s[0:i]
		for j := i; j < len(s); j += i {
			to := min(j+i, len(s))
			if s[j:to] != t {
				continue outer
			}
		}
		return true
	}
	return false
}


func part1(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		for i := r.from; i <= r.to; i += 1 {
			if isRepeated(strconv.Itoa(i)) {
				sum += i
			}
		}
	}
	return sum
}

func part2(ranges []Range) int {
	sum := 0
	for _, r := range ranges {
		for i := r.from; i <= r.to; i += 1 {
			if isRepeated2(strconv.Itoa(i)) {
				sum += i
			}
		}
	}
	return sum
}

func main() {
	ranges, err := readInput()
	if err != nil {
		panic(err)
	} 

	result1 := part1(ranges)
	fmt.Println("part 1 result: ", result1)

	result2 := part2(ranges)
	fmt.Println("part 2 result: ", result2)
}
