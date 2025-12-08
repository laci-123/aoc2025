package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	numbers []int
	operator byte
}

func readInput() ([]Problem, error) {
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

	problems := make([]Problem, 0)
	for i := 0; i < len(words[0]); i += 1 {
		numbers := make([]int, 0)
		var operator byte
		for j := 0; j < len(words); j += 1 {
			if j == len(words) - 1 {
				operator = words[len(words) - 1][i][0]
			} else {
				n, err := strconv.Atoi(words[j][i])
				if err != nil {
					return nil, err
				}
				numbers = append(numbers, n)
			}
		}
		problems = append(problems, Problem{numbers, operator})
	}
	return problems, nil
}

func part1(problems []Problem) int {
	sum := 0
	for _, problem := range problems {
		switch problem.operator {
		case '+':
			x := 0
			for _, n := range problem.numbers {
				x += n
			}
			sum += x
		case '*':
			x := 1
			for _, n := range problem.numbers {
				x *= n
			}
			sum += x
		}
	}
	return sum
}

func main() {
	problems, err := readInput()
	if err != nil {
		panic(err)
	} 

	result1 := part1(problems)
	fmt.Println("part 1 result: ", result1)
}
