package main

import (
	"fmt"

	"github.com/m7kvqbe1/advent-of-code-2023/pkg/day01"
	"github.com/m7kvqbe1/advent-of-code-2023/pkg/day02"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "advent-of-code-2023"}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "day01part1",
		Short: "Run Day 01 Part 1",
		Run: func(cmd *cobra.Command, args []string) {
			day01.Part1()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "day01part2",
		Short: "Run Day 01 Part 2",
		Run: func(cmd *cobra.Command, args []string) {
			day01.Part2()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "day02part1",
		Short: "Run Day 02 Part 1",
		Run: func(cmd *cobra.Command, args []string) {
			day02.Part1()
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "day02part2",
		Short: "Run Day 02 Part 2",
		Run: func(cmd *cobra.Command, args []string) {
			day02.Part2()
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
