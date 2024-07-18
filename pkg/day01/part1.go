package day01

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func Part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit := -1, -1

		// First
		for _, char := range line {
			if unicode.IsDigit(rune(char)) {
				firstDigit = int(char - '0')
				break
			}
		}

		// Last - iterate in reverse
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				lastDigit = int(line[i] - '0')
				break
			}
		}

		if firstDigit != -1 && lastDigit != -1 {
			calibrationValue := firstDigit*10 + lastDigit
			total += calibrationValue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	fmt.Println(total)
}
