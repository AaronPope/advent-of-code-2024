package main

import (
	"aoc2024/day1"
	"aoc2024/day2"
	"aoc2024/day3"
	"aoc2024/day4"
	utils "aoc2024/utils"
)

func main() {
	days := []utils.Problem{
		day4.Exec(),
		day3.Exec(),
		day2.Exec(),
		day1.Exec(),
	}

	for _, day := range days {
		day.PrintResults()
	}
}
