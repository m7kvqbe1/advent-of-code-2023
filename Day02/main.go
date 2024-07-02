package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseCounts(game string) (r []int, g []int, b []int) {
	re := regexp.MustCompile(`(\d+) (red|green|blue)`)

	for _, subset := range strings.Split(strings.Split(game, ": ")[1], "; ") {
		rCount, gCount, bCount := 0, 0, 0

		for _, match := range re.FindAllStringSubmatch(subset, -1) {
			count, _ := strconv.Atoi(match[1])
			switch match[2] {
			case "red":
				rCount += count
			case "green":
				gCount += count
			case "blue":
				bCount += count
			}
		}

		r = append(r, rCount)
		g = append(g, gCount)
		b = append(b, bCount)
	}

	return r, g, b
}

// func isGamePossible(game string, maxR, maxG, maxB int) bool {
// 	r, g, b := parseCounts(game)

// 	for _, red := range r {
// 		if red > maxRed {
// 			return false
// 		}
// 	}

// 	for _, green := range g {
// 		if green > maxGreen {
// 			return false
// 		}
// 	}

// 	for _, blue := range b {
// 		if blue > maxBlue {
// 			return false
// 		}
// 	}

// 	return true
// }

func minCubesRequired(game string) (int, int, int) {
	r, g, b := parseCounts(game)

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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer file.Close()

	// maxR, maxG, maxB := 12, 13, 14
	sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		game := scanner.Text()
		// re := regexp.MustCompile(`Game (\d+):`)
		// match := re.FindStringSubmatch(game)
		// gameID, _ := strconv.Atoi(match[1])

		// if isGamePossible(game, maxR, maxG, maxB) {
		// 	sum += gameID
		// }

		minR, minG, minB := minCubesRequired(game)

		power := minR * minG * minB
		sum += power
	}

	fmt.Printf("Sum of possible IDs: %d\n", sum)
}
