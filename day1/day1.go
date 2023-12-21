package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numMap = map[string]string{
	"zero":  "0",
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

func RunDay1() {
	var lines = readFile()
	var sum = 0

	for _, line := range lines {
		sum += day1Helper(line)
	}

	fmt.Println(sum)
}

func RunDay1Part2() {
	var lines = readFile()
	var sum = 0

	for _, line := range lines {
		sum += day1Part2Helper(line)
	}

	fmt.Println(sum)
}

func convertLineStrNum(line string) string {
	newStr := ""
	start := 0

	// really bad on^3 loop
	for key := range numMap {

		for {
			idx := strings.Index(line[start:], key)
            
		}
	}
}

func day1Part2Helper(line string) int {
	// TODO: convert line with above func
	// they should all be positive numbers
	var first rune = 0
	var second rune = 0

	for _, character := range line {
		if !unicode.IsDigit(character) {
			continue
		}

		if first == 0 {
			first = character
			continue
		}

		second = character
	}

	if second == 0 {
		second = first
	}

	result, err := strconv.Atoi(string(first) + string(second))
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func day1Helper(line string) int {
	// they should all be positive numbers
	var first rune = 0
	var second rune = 0

	for _, character := range line {
		if !unicode.IsDigit(character) {
			continue
		}

		if first == 0 {
			first = character
			continue
		}

		second = character
	}

	if second == 0 {
		second = first
	}

	result, err := strconv.Atoi(string(first) + string(second))
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func readFile() []string {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}
