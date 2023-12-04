package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.Open("input.txt")
	checkError(err)

	readFile := bufio.NewScanner(content)

	defer content.Close()

	var lines []string

	for readFile.Scan() {
		lines = append(lines, readFile.Text())
	}

	result := parseLines(lines)
	fmt.Println(result)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseLines(lines []string) int {
	var result int
	for _, line := range lines {
		card := strings.Split(line, ": ")
		numbers := strings.Split(card[1], " | ")

		wNumbers := strings.SplitAfter(numbers[0], " ")
		mNumbers := strings.SplitAfter(numbers[1], " ")

		winnerNumbers := convertToNumber(wNumbers)
		myNumbers := convertToNumber(mNumbers)

		var matches int
		points := 0

		for _, mNum := range myNumbers {
			for _, wNum := range winnerNumbers {
				if mNum == wNum {
					//fmt.Println("Match", mNum, wNum)
					matches++
					if matches >= 3 {
						points *= 2
					} else {
						points++
					}
					fmt.Println("Points", points)
				}
			}
		}

		result += points
	}

	return result
}

func convertToNumber(numbers []string) []int {
	var convertedNumbers []int
	for _, n := range numbers {
		num := strings.TrimSpace(n)
		ok, _ := isDigit(num)

		if ok {
			num, err := strconv.Atoi(num)
			checkError(err)
			convertedNumbers = append(convertedNumbers, num)
		}
	}
	return convertedNumbers
}

func isDigit(char string) (bool, error) {
	_, err := strconv.Atoi(char)
	if err != nil {
		return false, err
	}
	return true, nil
}
