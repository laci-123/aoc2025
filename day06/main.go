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

func main() {
	words, err := readInput()
	if err != nil {
		panic(err)
	} 

	result1, err := part1(words)
	if err != nil {
		panic(err)
	}
	fmt.Println("part 1 result: ", result1)
}
