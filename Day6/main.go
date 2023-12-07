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

	defer content.Close()

	var lines []string

	readFile := bufio.NewScanner(content)

	for readFile.Scan() {
		lines = append(lines, readFile.Text())
	}

	parseLines(lines)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func parseLines(lines []string) {
	var times []int
	var distances []int
	for _, line := range lines {
		lineValue := strings.Split(line, ": ")
		if lineValue[0] == "Time" {
			time := strings.Split(lineValue[1], " ")
			times = convertToInt(time)
		}

		if lineValue[0] == "Distance" {
			distance := strings.Split(lineValue[1], " ")
			distances = convertToInt(distance)
		}

	}

	result := 1
	for i := 0; i < len(times); i++ {
		result *= part1(times[i], distances[i])
	}

	fmt.Println(result)
}

func part1(time, distance int) int {
	result := 0
	for i := 1; i < time; i++ {
		if ((time - i) * i) > distance {
			result++
		}
	}
	return result
}

func convertToInt(values []string) []int {
	var numValue []int
	for _, value := range values {
		if len(value) == 0 {
			continue
		}
		num := strings.TrimSpace(value)
		numConverted, err := strconv.Atoi(num)
		checkError(err)
		numValue = append(numValue, numConverted)
	}
	return numValue
}
