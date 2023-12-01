package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)


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

        var result int = 0

        for _, line := range lines {
                number := regexp.MustCompile("[0-9]")
                matches := number.FindAllString(line, -1)

                valueCombine := matches[0] + matches[len(matches) -1]
                value, err := strconv.Atoi(valueCombine)

                if err != nil {
                        panic(err)
                }

                result = finalResult(&result, value)
        }

        fmt.Println(result)
}

func finalResult(result *int, value int) int {
        *result += value
        return *result
}
