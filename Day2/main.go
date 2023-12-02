package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var setsCubes = map[string]int{}

var maxCount = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	content, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer content.Close()

	readeFile := bufio.NewScanner(content)
	readeFile.Split(bufio.ScanLines)
	var lines []string

	for readeFile.Scan() {
		lines = append(lines, readeFile.Text())
	}

	partOne := partOne(lines)
	fmt.Println(partOne)

	partTwo := partTwo(lines)
	fmt.Println(partTwo)
}

func partOne(lines []string) int {
	var total int = 0
	for idx, line := range lines {
		values, isValid := getValidSets(line, idx+1)
		if isValid {
			total += values
		}
	}

	return total
}

func getValidSets(line string, idx int) (int, bool) {
	game := strings.Split(line, ": ")
	sets := strings.Split(game[1], "; ")

	for _, set := range sets {
		cubes := strings.Split(set, ", ")

		for _, color := range cubes {
			cube := strings.Split(color, " ")
			value, err := strconv.Atoi(cube[0])
			if err != nil {
				panic(err)
			}

			setsCubes[cube[1]] = value
			if setsCubes[cube[1]] > maxCount[cube[1]] {
				return idx, false
			}
		}

	}

	return idx, true
}

func partTwo(lines []string) int {
	total := 0

	for _, line := range lines {
		var result = map[string]int{}
		getPower(line, &result)
		total += result["red"] * result["green"] * result["blue"]
	}

	return total
}

func getPower(line string, result *map[string]int) *map[string]int {
	game := strings.Split(line, ": ")
	sets := strings.Split(game[1], "; ")

	for _, set := range sets {
		cubes := strings.Split(set, ", ")

		for _, color := range cubes {
			cube := strings.Split(color, " ")
			value, err := strconv.Atoi(cube[0])
			if err != nil {
				panic(err)
			}

			if (*result)[cube[1]] < value {
				(*result)[cube[1]] = value
			}
		}
	}
	return result
}
