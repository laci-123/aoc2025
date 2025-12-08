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


func main() {
	db, err := readInput()
	if err != nil {
		panic(err)
	} 

	result1 := part1(db)
	fmt.Println("part 1 result: ", result1)
}
