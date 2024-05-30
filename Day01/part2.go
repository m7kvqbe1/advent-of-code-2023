package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	wordToDigit := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)

	total := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllString(line, -1)

		if len(matches) < 2 {
			continue
		}

		first := matches[0]
		last := matches[len(matches)-1]

		var firstDigit, lastDigit int
		if unicode.IsDigit(rune(first[0])) {
			firstDigit = int(first[0] - '0')
		} else {
			firstDigit = wordToDigit[first]
		}

		if unicode.IsDigit(rune(last[0])) {
			lastDigit = int(last[0] - '0')
		} else {
			lastDigit = wordToDigit[last]
		}

		calibrationValue := firstDigit*10 + lastDigit
		total += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	fmt.Println(total)
}
