package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type IdRange struct {
	begin int
	end int
}

type Database struct{
	ranges []IdRange
	ids []int
}

func readInput() (Database, error) {
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		return Database{ranges: nil, ids: nil}, err
	}
	ranges := make([]IdRange, 0)
	ids := make([]int, 0)
	firstPart := true
	for line := range strings.Lines(string(content)) {
		trimmedLine := strings.Trim(line, " \n")
		if trimmedLine == "" {
			firstPart = false
			continue
		}
		if firstPart {
			var begin, end int
			_, err := fmt.Sscanf(trimmedLine, "%v-%v\n", &begin, &end)
			if err != nil {
				return Database{ranges: nil, ids: nil}, err
			}
			ranges = append(ranges, IdRange{begin, end})
		} else {
			var id int
			_, err := fmt.Sscanf(trimmedLine, "%v\n", &id)
			if err != nil {
				return Database{ranges: nil, ids: nil}, err
			}
			ids = append(ids, id)
		}
	}
	return Database{ranges, ids}, nil
}

func part1(db Database) int {
	sum := 0
	for _, id := range db.ids {
		for _, rng := range db.ranges {
			if rng.begin <= id && id <= rng.end {
				sum += 1
				break
			}
		}
	}
	return sum
}

func rangeConcat(rng1, rng2 IdRange) (bool, IdRange) {
	if rng1.end < rng2.begin || rng2.end < rng1.begin {
		return false, IdRange{-1, -1}
	}
	if rng1.begin == rng2.begin && rng1.end == rng2.end {
		return false, IdRange{-1, -1}
	}
	return true, IdRange{begin: min(rng1.begin, rng2.begin), end: max(rng1.end, rng2.end)}
}

func rangeConcatMany(rngs map[IdRange]bool) map[IdRange]bool {
	newRngs := make(map[IdRange]bool)
	concatHappend := false
	for rng1, b1 := range rngs {
		if !b1 {
			continue
		}
		for rng2, b2 := range rngs {
			if !b2 {
				continue
			}
			ok, newRng := rangeConcat(rng1, rng2)
			if ok {
				newRngs[newRng] = true
				concatHappend = true
				if !(rng1.begin == newRng.begin && rng1.end == newRng.end) {
					newRngs[rng1] = false
				}
				if !(rng2.begin == newRng.begin && rng2.end == newRng.end) {
					newRngs[rng2] = false
				}
			} else {
				if _, ok := newRngs[rng1]; !ok {
					newRngs[rng1] = true
				}
				if _, ok := newRngs[rng2]; !ok {
					newRngs[rng2] = true
				}
			}
		}
	}
	if concatHappend {
		return rangeConcatMany(newRngs)
	} else {
		return newRngs
	}
}

func part2(db Database) int {
	rngs := make(map[IdRange]bool)
	for _, rng := range db.ranges {
		rngs[rng] = true
	}
	newRngs := rangeConcatMany(rngs)
	sum := 0
	for rng, b := range newRngs {
		if !b {
			continue
		}
		sum += rng.end - rng.begin + 1
	}
	return sum
}


func main() {
	db, err := readInput()
	if err != nil {
		panic(err)
	} 

	result1 := part1(db)
	fmt.Println("part 1 result: ", result1)
	result2 := part2(db)
	fmt.Println("part 2 result: ", result2)
}
