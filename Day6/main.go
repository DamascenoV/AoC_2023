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
	var timeValue int
	var distanceValue int
	for _, line := range lines {
		lineValue := strings.Split(line, ": ")
		if lineValue[0] == "Time" {
			time := strings.Split(lineValue[1], " ")
			time2 := part2Parser(time)
			times = convertToInt(time)
			tValue, err := strconv.Atoi(time2)
			checkError(err)
			timeValue = tValue
		}

		if lineValue[0] == "Distance" {
			distance := strings.Split(lineValue[1], " ")
			distance2 := part2Parser(distance)
			distances = convertToInt(distance)
			dValue, err := strconv.Atoi(distance2)
			checkError(err)
			distanceValue = dValue
		}
	}

	resultPart1 := 1
	for i := 0; i < len(times); i++ {
		resultPart1 *= getResult(times[i], distances[i])
	}

	resultPart2 := getResult(timeValue, distanceValue)

	fmt.Println("Part 1: ", resultPart1)
	fmt.Println("Part 2: ", resultPart2)

}

func getResult(time, distance int) int {
	result := 0
	for i := 1; i < time; i++ {
		if ((time - i) * i) > distance {
			result++
		}
	}
	return result
}

func part2Parser(values []string) string {
	valueConcat := strings.Join(values, "")
	value := strings.TrimSpace(valueConcat)

	return value
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
