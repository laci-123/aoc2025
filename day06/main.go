package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readInput() ([][]string, error) {
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}

	words := make([][]string, 0)
	for line := range strings.Lines(string(content)) {
		wordsLine := make([]string, 0)
		var word strings.Builder
		for _, c := range line {
			if c == ' ' || c == '\n' {
				if word.Len() > 0 {
					wordsLine = append(wordsLine, word.String())
					word.Reset()
				} else {
					continue
				}
			} else {
				word.WriteRune(c)
			}
		}
		words = append(words, wordsLine)
	}

	return words, nil
}

func readInput2() ([][]rune, error) {
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	runes := make([][]rune, 0)
	for line := range strings.Lines(string(content)) {
		cLine := make([]rune, 0)
		for _, c := range line {
			cLine = append(cLine, c)
		}
		runes = append(runes, cLine)
	}
	return runes, nil
}

func part1(words [][]string) (int, error) {
	sum := 0
	for i := 0; i < len(words[0]); i += 1 {
		var operator byte
		x := 0
		for j := len(words)-1; j >= 0; j -= 1 {
			if j == len(words) - 1 {
				operator = words[len(words) - 1][i][0]
			} else {
				n, err := strconv.Atoi(words[j][i])
				if err != nil {
					return 0, err
				}
				switch operator {
				case '+':
					x += n
				case '*':
					if x == 0 {
						x = 1
					}
					x *= n
				}
			}
		}
		sum += x
	}
	return sum, nil
}

func part2(runes [][]rune) (int, error) {
	sum := 0
	columns := make([]int, 0)
	var op rune
	for x := len(runes[0])-1; x >= 0; x -= 1 {
		n := 0
		p := 1
		allSpace := true
		for y := len(runes)-1; y >= 0; y -= 1 {
			c := runes[y][x]
			switch c {
			case '+':
				allSpace = false
				op = '+'
			case '*':
				allSpace = false
				op = '*'
			case ' ', '\n':
				continue
			default:
				allSpace = false
				n += int(c - '0') * p
				p *= 10
			}
		}
		if !allSpace || x == 0 {
			columns = append(columns, n)
		}
		if allSpace || x == 0 {
			var x int
			switch op {
			case '+':
				x = 0
			case '*':
				x = 1
			}
			for _, n := range columns {
				switch op {
				case '+':
					x += n
				case '*':
					x *= n
				}
			}
			sum += x
			columns = columns[:0]
		}
	}
	return sum, nil
}

func main() {
	// words, err := readInput()
	// if err != nil {
	// 	panic(err)
	// } 
	// result1, err := part1(words)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("part 1 result: ", result1)

	runes, err := readInput2()
	if err != nil {
		panic(err)
	}

	result2, err := part2(runes)
	if err != nil {
		panic(err)
	}
	fmt.Println("part 2 result: ", result2)
}
