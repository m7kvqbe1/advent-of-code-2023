package day02

import (
	"bufio"
	"fmt"
	"os"
)

func minCubesRequired(game string) (int, int, int) {
	r, g, b := ParseCounts(game)

	max := func(counts []int) int {
		max := 0

		for _, count := range counts {
			if count > max {
				max = count
			}
		}

		return max
	}

	return max(r), max(g), max(b)
}

func Part2() {
	file, err := os.Open("pkg/day02/input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		game := scanner.Text()

		minR, minG, minB := minCubesRequired(game)

		power := minR * minG * minB
		sum += power
	}

	fmt.Printf("Sum of all powers: %d\n", sum)
}
