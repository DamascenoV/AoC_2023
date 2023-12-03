package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Number struct {
	value     string
	startChar int
	lastChat  int
	line      int
}

type Symbol struct {
	line     int
	position int
}

func main() {
	content, err := os.Open("input.txt")
	checkError(err)

	defer content.Close()

	readFile := bufio.NewScanner(content)

	var lines []string
	for readFile.Scan() {
		lines = append(lines, readFile.Text())
	}

	numbers, symbols := parseLine(lines)

	result := calculateResult(numbers, symbols)

	fmt.Println(result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseLine(lines []string) ([]*Number, []*Symbol) {
	var digit string
	var firstChar int
	var numbers []*Number
	var symbols []*Symbol

	for idx, line := range lines {
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if _, ok := isDigit(char); ok {
				digit += char
				if firstChar == 0 {
					firstChar = i
				}
			} else {
				if digit != "" {
					numbers = append(
						numbers,
						&Number{
							value:     digit,
							startChar: firstChar,
							lastChat:  i - 1,
							line:      idx,
						},
					)
					digit = ""
					firstChar = 0
				}

				if char != "." {
					symbols = append(
						symbols,
						&Symbol{
							line:     idx,
							position: i,
						},
					)
				}
			}
		}

		if digit != "" {
			numbers = append(
				numbers,
				&Number{
					value:     digit,
					startChar: firstChar,
					lastChat:  len(line) - 1,
					line:      idx,
				},
			)
			digit = ""
			firstChar = 0
		}
	}

	return numbers, symbols
}

func calculateResult(numbers []*Number, symbols []*Symbol) int {
	var result int
	for _, number := range numbers {
		var isAdjecent bool
		for _, symbol := range symbols {
			if checkAdjecent(number, symbol) {
				for i := number.startChar - 1; i <= number.lastChat+1; i++ {
					if i == symbol.position {
						isAdjecent = true

						number, err := strconv.Atoi(number.value)
						checkError(err)

						result += number
						break
					}
				}
				if isAdjecent {
					break
				}
			}
		}
	}

	return result
}

func checkAdjecent(number *Number, symbol *Symbol) bool {
	return number.line == symbol.line || number.line+1 == symbol.line || number.line-1 == symbol.line
}

func isDigit(value string) (int, bool) {
	n, err := strconv.Atoi(value)
	if err != nil {
		return 0, false
	}

	return n, true
}
