package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	content, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	readFile := bufio.NewScanner(content)
	readFile.Split(bufio.ScanLines)
	var lines []string

	for readFile.Scan() {
		lines = append(lines, readFile.Text())
	}

	content.Close()

	regexPartOne := "[1-9]"
	regexPartTwo := "(one|two|three|four|five|six|seven|eight|nine|[1-9])"

	partOne := result(lines, regexPartOne)
	partTwo := result(lines, regexPartTwo)

	fmt.Println(partOne)
	fmt.Println(partTwo)
}

func result(lines []string, regex string) int {
	var result int = 0

	for _, line := range lines {
		number := regexp.MustCompile(regex)
		matches := number.FindAllString(line, -1)

		if len(matches) == 0 {
			continue
		}

		valueCombine := convertSpelledNumber(matches[0]) + convertSpelledNumber(matches[len(matches)-1])
		value, err := strconv.Atoi(valueCombine)

		if err != nil {
			panic(err)
		}

		result = finalResult(&result, value)
	}

	return result
}

func finalResult(result *int, value int) int {
	*result += value
	return *result
}

func convertSpelledNumber(number string) string {
	if _, ok := numbers[number]; ok {
		return numbers[number]
	}

	return number
}
